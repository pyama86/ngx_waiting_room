package waitingroom

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/securecookie"
)

func TestAccessController_setAllowedNo(t *testing.T) {
	tests := []struct {
		name    string
		domain  string
		want    int64
		wantErr bool
	}{
		{
			name:   "ok",
			domain: string(securecookie.GenerateRandomKey(64)),
			want:   1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisClient := redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
			})
			a := &AccessController{
				QueueBase: QueueBase{
					config: &Config{
						AllowUnitNumber: 1000,
						QueueEnableSec:  600,
					},
					redisClient: redisClient,
					cache:       NewCache(redisClient, &Config{}),
				},
			}
			got, err := a.setAllowedNo(context.Background(), tt.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessController.setAllowedNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccessController.setAllowedNo() = %v, want %v", got, tt.want)
			}

			v, err := redisClient.Get(context.Background(), tt.domain+"_allow_no").Result()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessController.setAllowedNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if v != strconv.Itoa(int(tt.want)) {
				t.Errorf("AccessController.setAllowedNo() = %v, want %v", v, tt.want)
			}

			tv, err := redisClient.TTL(context.Background(), tt.domain+"_allow_no").Result()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessController.setAllowedNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tv != 600*time.Second {
				t.Errorf("AccessController.setAllowedNo() = %v, want %v", tv, 600)
			}
		})
	}
}

func TestAccessController_Do(t *testing.T) {
	tests := []struct {
		ctx        context.Context
		beforeHook func(string, *redis.Client)
		name       string
		domain     string
		want       int
		wantErr    bool
	}{
		{
			name:   "ok",
			domain: string(securecookie.GenerateRandomKey(12)),
			beforeHook: func(key string, redisClient *redis.Client) {
				redisClient.Del(context.Background(), enableDomainKey)
				redisClient.SAdd(context.Background(), enableDomainKey, key)
				redisClient.SetEX(context.Background(), key+"_allow_no", "1", 3*time.Second)
			},
			want: 1001,
		},
		{
			name:   "skip update because not target domain",
			domain: string(securecookie.GenerateRandomKey(12)),
			beforeHook: func(key string, redisClient *redis.Client) {
				redisClient.Del(context.Background(), enableDomainKey)
				redisClient.SAdd(context.Background(), enableDomainKey, "hoge")
			},
			want: 0,
		},
		{
			name:   "skip update because not enabled domain",
			domain: string(securecookie.GenerateRandomKey(12)),
			beforeHook: func(key string, redisClient *redis.Client) {
				redisClient.Del(context.Background(), enableDomainKey)
				redisClient.SAdd(context.Background(), enableDomainKey, key)
			},
			want: 0,
		},
		{
			name:   "skip update because not before update time",
			domain: string(securecookie.GenerateRandomKey(12)),
			beforeHook: func(key string, redisClient *redis.Client) {
				redisClient.Del(context.Background(), enableDomainKey)
				redisClient.SAdd(context.Background(), enableDomainKey, key)
				redisClient.SetEX(context.Background(), key+"_allow_no", "1", 600*time.Second)
				redisClient.SetEX(context.Background(), key+"_lock_allow_no", "1", 10*time.Second)
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisClient := redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
			})
			a := &AccessController{
				QueueBase: QueueBase{
					config: &Config{
						AllowUnitNumber: 1000,
						QueueEnableSec:  600,
					},
					redisClient: redisClient,
					cache:       NewCache(redisClient, &Config{}),
				},
			}

			if tt.beforeHook != nil {
				tt.beforeHook(tt.domain, redisClient)
			}
			if err := a.Do(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("AccessController.Do() error = %v, wantErr %v", err, tt.wantErr)
			}

			v, err := redisClient.Get(context.Background(), tt.domain+"_allow_no").Result()
			if tt.want != 0 {
				if (err != nil) != tt.wantErr {
					t.Errorf("AccessController.setAllowedNo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if v != strconv.Itoa(int(tt.want)) {
					t.Errorf("AccessController.setAllowedNo() = %v, want %v", v, tt.want)
				}

				tv, err := redisClient.TTL(context.Background(), tt.domain+"_allow_no").Result()
				if (err != nil) != tt.wantErr {
					t.Errorf("AccessController.Do() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if tv != 600*time.Second {
					t.Errorf("AccessController.Do() = %v, want %v", tv, 600)
				}
			} else {
				if err != redis.Nil {
					t.Errorf("AccessController.setAllowedNo() value = %v error = %v", v, err)
					return
				}
			}
		})
	}
}