module github.com/aliasliao/shadow

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/protobuf v1.25.0
	v2ray.com/core v0.0.0-00010101000000-000000000000
)

replace v2ray.com/core v0.0.0-00010101000000-000000000000 => github.com/aliasliao/v2ray-core v1.24.5-0.20200903151501-dc61c1a1b5a9
