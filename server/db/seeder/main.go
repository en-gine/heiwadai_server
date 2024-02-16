package main

import (
	"flag"
)

func main() {
	runStoreSeeder := flag.Bool("StoreSeeder", false, "StoreSeederを呼び出す")
	runUserSeeder := flag.Bool("UserSeeder", false, "UserSeederを呼び出す")
	runAdminSeeder := flag.Bool("AdminSeeder", false, "AdminSeederを呼び出す")
	runCouponSeeder := flag.Bool("CouponSeeder", false, "CouponSeederを呼び出す")
	runOnlyServer := flag.Bool("OnlyServer", false, "サーバーのみ起動")
	flag.Parse()

	if *runStoreSeeder {
		StoreSeeder()
	}

	if *runUserSeeder {
		UserSeeder()
	}

	if *runAdminSeeder {
		AdminSeeder()
	}

	if *runCouponSeeder {
		CouponSeeder()
	}

	if *runOnlyServer {
		OnlyServer()
	}
}
