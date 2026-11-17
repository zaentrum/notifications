// Package config centralises environment-driven configuration for the
// notifications service.
package config

import (
	"os"
	"strings"
)

// Config holds all runtime tunables.
type Config struct {
	ListenAddr    string
	KafkaBrokers  []string
	KafkaGroupID  string
	KafkaTopics   []string
	APNsKeyPath   string
	APNsTeamID    string
	APNsKeyID     string
	FCMCredsPath  string
	WebPushSubKey string
}

// FromEnv builds a Config from process environment.
func FromEnv() Config {
	return Config{
		ListenAddr:    envOr("LISTEN_ADDR", ":8080"),
		KafkaBrokers:  splitCSV(envOr("KAFKA_BROKERS", "kafka:9092")),
		KafkaGroupID:  envOr("KAFKA_GROUP_ID", "notifications"),
		KafkaTopics:   splitCSV(envOr("KAFKA_TOPICS", "stube.notify.user.push")),
		APNsKeyPath:   os.Getenv("APNS_KEY_PATH"),
		APNsTeamID:    os.Getenv("APNS_TEAM_ID"),
		APNsKeyID:     os.Getenv("APNS_KEY_ID"),
		FCMCredsPath:  os.Getenv("FCM_CREDS_PATH"),
		WebPushSubKey: os.Getenv("WEBPUSH_SUB_KEY"),
	}
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
