package main

import (
    "context"
    "github.com/redis/go-redis/v9"
    "fmt"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "polly.iran.liara.ir:32744",
        Password: "1NJLiuSk1zUjiYTLgmZqzKqR",
        DB:       0,  
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}

/*
! THIS DATA IS NOT WORK DO NOT TEST.
host: polly.iran.liara.ir
port: 32744
pass: 1NJLiuSk1zUjiYTLgmZqzKqR
uri: redis://:1NJLiuSk1zUjiYTLgmZqzKqR@polly.iran.liara.ir:32744/0
*/