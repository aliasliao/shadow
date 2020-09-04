module github.com/aliasliao/shadow

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/protobuf v1.25.0
	v2ray.com/core v1.24.5-0.20200903151501-dc61c1a1b5a9
)

replace v2ray.com/core v1.24.5-0.20200903151501-dc61c1a1b5a9 => github.com/aliasliao/v2ray-core v1.24.5-0.20200903151501-dc61c1a1b5a9
