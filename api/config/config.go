package config

import (
	"context"
	"errors"
	"os"
	"reflect"

	"github.com/glassonion1/logz"
)

type Config struct {
	Port       string `env:"API_PORT"`
	ProjectID  string `env:"PROJECT_ID"`
	InstanceID string `env:"INSTANCE_ID"`
	DatabaseID string `env:"DB_ID"`
	CF         string `env:"CF"`
	Stage      string `env:"STAGE"`
}

func NewConfig() *Config {
	config := new(Config)

	logz.Infof(context.Background(), "env loaded...")
	err := setEnv(config)
	if err != nil {
		panic(err)
	}

	return config
}

// 環境変数をパース
func setEnv(config *Config) error {
	ctx := context.Background()
	cType := reflect.TypeOf(config).Elem()
	cValue := reflect.ValueOf(config).Elem()

	err := parseEnv(ctx, cType, cValue)
	if err != nil {
		return err
	}

	return nil
}

// modeが"secret"の場合に処理を分ける
func parseEnv(ctx context.Context, cType reflect.Type, cValue reflect.Value) error {
	// cType と cValue は Config 構造体の型と値をリフレクト
	for i := 0; i < cType.NumField(); i++ {
		valueField := cValue.Field(i)
		typeField := cType.Field(i)

		// フィールドがさらに構造体である場合（ネストされた構造体）
		// その構造体に対しても同じ環境変数の設定処理（parseEnv関数）を再帰的
		if valueField.Kind() == reflect.Struct {
			cc := valueField.Addr().Interface()
			ccType := reflect.TypeOf(cc).Elem()
			ccValue := reflect.ValueOf(cc).Elem()
			err := parseEnv(ctx, ccType, ccValue)
			if err != nil {
				return err
			}
			continue
		}

		// envタグから環境変数の名前を取得
		envName := typeField.Tag.Get("env")
		print(envName)
		if envName == "" {
			continue
		}

		// configのフィールドに値をセットできない場合はエラーを出す
		if !valueField.CanSet() {
			return ErrCannotSetEnvValue
		}

		var envValue string
		// ローカル環境ではない、かつmodeがSecretの時はSecretManagerから値を取得する
		if os.Getenv("STAGE") != "local" && typeField.Tag.Get("mode") == "secret" {
			// TODO:secret managerの実装
			logz.Errorf(ctx, "secret manager is not definition")
			// value, err := secret.GetSecret(ctx, os.Getenv(envName))
			// if err == nil {
			// envValue = value
			// } else {
			// logz.Errorf(ctx, "failed to get secret: %v", err)
			// }
		} else {
			// ローカルは.envから取得
			envValue = os.Getenv(envName)
			print(envValue)
		}

		// ローカル環境ではない、かつmodeがSecretの時はSecretManagerから値を取得する
		if os.Getenv("STAGE") == "local" && typeField.Tag.Get("mode") != "secret" {
			// ローカルは.envから取得
			envValue = os.Getenv(envName)
			print(envValue)
		}

		// 環境変数が設定されていない、かつそのフィールドが秘密情報でない場合は、デフォルト値を設定します。
		// if envValue == "" && typeField.Tag.Get("mode") != "secret" {
		// envValue = typeField.Tag.Get("default")
		// }

		valueField.SetString(envValue)
	}

	return nil
}

var (
	// ErrCannotSetEnvValue is the error because of config fields is not public
	ErrCannotSetEnvValue = errors.New("cannot set env value")
)
