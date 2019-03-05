## Golang TimeZone

### Config.yaml

```bash
timezone: Asia/Hong_Kong
```

### load time zone

```go
time.Location

// 1. via IANA Database
time.LoadLocation("location")
// 2. via offset
time.FixedZone("UTC-8", -8 **60 **60)

```

### converge 

```go
tl := time.UTC
time.Now().In(tl)
_ = t1
```

### modify

```go
var time1 time.Time
var time2 time.Time
var time1Zone *time.Location
// via time.Date -> time struct instance
time1 = time.Date(time2.Year(), time2.Month(),
	0, 0, 0, 0, 0, time1Zone)
_ = time1
```

### local mean time (LMT)

    香港日佔時期，又稱為香港日治時期或香港淪陷時期，是指第二次世界大戰時
    大日本帝國軍事占領香港的時期：由1941年12月25日香港總督楊慕琦投降起，
    至1945年8月15日日本無條件投降為止；香港人俗稱這段時期為「三年零八個月」