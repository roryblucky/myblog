package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	"myblog/logger"
)

var cc cache.Cache

func init() {
	cache, err := cache.NewCache("memory", `{interval:300}`)
	if err != nil {
		logger.Error("Cache init failed, error msg: %s", err.Error())
		panic(err.Error())
	}
	cc = cache
}

func SetCache(key string, data interface{}, timeout int64) error {
	if IsEmpty(key) {
		return errors.New("cache key cannot be null")
	}

	b, err := encodeData(data)
	if err != nil {
		return err
	}

	err = cc.Put(key, b, timeout)

	if err != nil {
		logger.Error("cache %s failed", key)
		return err
	}
	return nil
}

func DelCache(key string) error {
	if IsEmpty(key) {
		return errors.New("cache key cannot be null")
	}

	err := cc.Delete(key)

	if err != nil {
		logger.Warn("delete cache %s failed, maybe not exist", key)
		return err
	}
	return nil
}

func GetDataFromCache(key string, target interface{}) error {
	if IsEmpty(key) {
		return errors.New("cache key cannot be null")
	}

	b := cc.Get(key)
	if b == nil {
		return errors.New(fmt.Sprintf("%s cache is not exist", key))
	}
	err := decodeData(b.([]byte), target)

	if err != nil {
		logger.Error("get cache %s failed", key)
		return err
	}
	return nil
}

func encodeData(data interface{}) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buff)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func decodeData(data []byte, target interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(target)
}
