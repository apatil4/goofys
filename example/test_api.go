package main

import (
	goofys "github.com/apatil4/goofys/api"
	common "github.com/apatil4/goofys/api/common"

	"fmt"
	"context"
)

func main() {
	config := common.FlagStorage{
		MountPoint: "/tmp/s3",
		DirMode:    0755,
		FileMode:   0644,
	}

	_, mp, err := goofys.Mount(context.Background(), "goofys", &config)
	if err != nil {
		panic(fmt.Sprintf("Unable to mount %v: %v", config.MountPoint, err))
	} else {
		mp.Join(context.Background())
	}
}
