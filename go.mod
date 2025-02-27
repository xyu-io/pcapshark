module github.com/xyu-io/pcapshark

go 1.22.0

toolchain go1.22.8

// fork of github.com/itchyny/gojq, see github.com/wader/gojq fq branch
require github.com/wader/gojq v0.12.1-0.20250208151254-0aa7b87b2c2b

require (
	// bump: gomod-creasty-defaults /github\.com\/creasty\/defaults v(.*)/ https://github.com/creasty/defaults.git|^1
	// bump: gomod-creasty-defaults command go get github.com/creasty/defaults@v$LATEST && go mod tidy
	// bump: gomod-creasty-defaults link "Source diff $CURRENT..$LATEST" https://github.com/creasty/defaults/compare/v$CURRENT..v$LATEST
	github.com/creasty/defaults v1.8.0

	// bump: gomod-gopacket /github\.com\/gopacket\/gopacket v(.*)/ https://github.com/gopacket/gopacket.git|^1
	// bump: gomod-gopacket command go get github.com/gopacket/gopacket@v$LATEST && go mod tidy
	// bump: gomod-gopacket link "Release notes" https://github.com/gopacket/gopacket/releases/tag/v$LATEST
	github.com/gopacket/gopacket v1.3.1

	// bump: gomod-copystructure /github\.com\/mitchellh\/copystructure v(.*)/ https://github.com/mitchellh/copystructure.git|^1
	// bump: gomod-copystructure command go get github.com/mitchellh/copystructure@v$LATEST && go mod tidy
	// bump: gomod-copystructure link "CHANGELOG" https://github.com/mitchellh/copystructure/blob/master/CHANGELOG.md
	github.com/mitchellh/copystructure v1.2.0

	// bump: gomod-mapstructure /github\.com\/mitchellh\/mapstructure v(.*)/ https://github.com/mitchellh/mapstructure.git|^1
	// bump: gomod-mapstructure command go get github.com/mitchellh/mapstructure@v$LATEST && go mod tidy
	// bump: gomod-mapstructure link "CHANGELOG" https://github.com/mitchellh/mapstructure/blob/master/CHANGELOG.md
	github.com/mitchellh/mapstructure v1.5.0

	// bump: gomod-golang-x-crypto /golang\.org\/x\/crypto v(.*)/ https://github.com/golang/crypto.git|^0
	// bump: gomod-golang-x-crypto command go get golang.org/x/crypto@v$LATEST && go mod tidy
	// bump: gomod-golang-x-crypto link "Tags" https://github.com/golang/crypto/tags
	golang.org/x/crypto v0.33.0

	// bump: gomod-golang-x-net /golang\.org\/x\/net v(.*)/ https://github.com/golang/net.git|^0
	// bump: gomod-golang-x-net command go get golang.org/x/net@v$LATEST && go mod tidy
	// bump: gomod-golang-x-net link "Tags" https://github.com/golang/net/tags
	golang.org/x/net v0.35.0

	// bump: gomod-golang/text /golang\.org\/x\/text v(.*)/ https://github.com/golang/text.git|^0
	// bump: gomod-golang/text command go get golang.org/x/text@v$LATEST && go mod tidy
	// bump: gomod-golang/text link "Source diff $CURRENT..$LATEST" https://github.com/golang/text/compare/v$CURRENT..v$LATEST
	golang.org/x/text v0.22.0

	// bump: gomod-gopkg.in/yaml.v3 /gopkg\.in\/yaml\.v3 v(.*)/ https://github.com/go-yaml/yaml.git|^3
	// bump: gomod-gopkg.in/yaml.v3 command go get gopkg.in/yaml.v3@v$LATEST && go mod tidy
	// bump: gomod-gopkg.in/yaml.v3 link "Source diff $CURRENT..$LATEST" https://github.com/go-yaml/yaml/compare/v$CURRENT..v$LATEST
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/gomarkdown/markdown v0.0.0-20250207164621-7a1f277a159e
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/itchyny/timefmt-go v0.1.6 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
