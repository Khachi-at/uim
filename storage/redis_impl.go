package storage

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/uim"
	"github.com/uim/wire/pkt"
	"google.golang.org/protobuf/proto"
)

const (
	LocationExpired = time.Hour * 48
)

type RedisStorage struct {
	cli *redis.Client
}

var _ uim.SessionStorage = (*RedisStorage)(nil)

func NewRedisStorage(cli *redis.Client) uim.SessionStorage {
	return &RedisStorage{
		cli: cli,
	}
}

func (r *RedisStorage) Add(session *pkt.Session) error {
	// save uim.Location
	loc := uim.Location{
		ChannelId: session.ChannelId,
		GateId:    session.GateId,
	}
	locKey := KeyLocation(session.Account, "")
	err := r.cli.Set(locKey, loc.Bytes(), LocationExpired).Err()
	if err != nil {
		return err
	}
	// save session
	snKey := KeySession(session.ChannelId)
	buf, _ := proto.Marshal(session)
	return r.cli.Set(snKey, buf, LocationExpired).Err()
}

// Delete s session.
func (r *RedisStorage) Delete(account, channelId string) error {
	locKey := KeyLocation(account, "")
	if err := r.cli.Del(locKey).Err(); err != nil {
		return err
	}
	snKey := KeySession(channelId)
	return r.cli.Del(snKey).Err()
}

// Get session by sessionID.
func (r *RedisStorage) Get(channelId string) (*pkt.Session, error) {
	snKey := KeySession(channelId)
	bts, err := r.cli.Get(snKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, uim.ErrSessionNil
		}
		return nil, err
	}
	var session pkt.Session
	_ = proto.Unmarshal(bts, &session)
	return &session, nil
}

func (r *RedisStorage) GetLocations(accounts ...string) ([]*uim.Location, error) {
	keys := KeyLocations(accounts...)
	list, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}
	var result = make([]*uim.Location, 0)
	for _, l := range list {
		if l == nil {
			continue
		}
		var loc uim.Location
		_ = loc.Unmarshal([]byte(l.(string)))
		result = append(result, &loc)
	}
	if len(result) == 0 {
		return nil, uim.ErrSessionNil
	}
	return result, nil
}

func (r *RedisStorage) GetLocation(account, device string) (*uim.Location, error) {
	key := KeyLocation(account, device)
	bts, err := r.cli.Get(key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, uim.ErrSessionNil
		}
		return nil, err
	}
	var loc uim.Location
	_ = loc.Unmarshal(bts)
	return &loc, nil
}

func KeySession(channel string) string {
	return fmt.Sprintf("login:sn:%s", channel)
}

func KeyLocation(account, device string) string {
	if device == "" {
		return fmt.Sprintf("login:loc:%s", account)
	}

	return fmt.Sprintf("login:loc:%s:%s", account, device)
}

func KeyLocations(accounts ...string) []string {
	arr := make([]string, len(accounts))
	for i, account := range accounts {
		arr[i] = KeyLocation(account, "")
	}
	return arr
}
