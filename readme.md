# Simple Crud :
[![Build & Test](https://github.com/MRdyRy/go-crud-food/actions/workflows/pipeline.yml/badge.svg?branch=master)](https://github.com/MRdyRy/go-crud-food/actions/workflows/pipeline.yml)
- <a href="https://go.dev/">go</a> 
- <a href="https://gin-gonic.com/">gin</a>
- <a href="https://gorm.io/index.html">gorm</a>
- <a href="https://www.postgresql.org/">postgresql</a>
- <a href="https://www.docker.com/">docker</a>
- <a href="https://www.gnu.org/software/make/manual/make.html">makefile</a>

# common rules in go :
- hanya 1 main func (jika > 1 maka akan menyebabkan error ketika compile)
- keyword berawalan huruf kecil = private
- keyword berawalan huruf besar = public
- keyword var bersifat optional
- untuk membuat var tanpa keyword var gunakan keyword :=
> ex: <br/>
var name <br/>
name := "ryan"
- type data :
> int8, int16, int32, int64 <br/>
float32, float64
- type data alias :
> rune = int32<br/>
byte = uint8<br/>
int = Minimal int32<br/>
uint = Minimal int32<br/>
- untuk deklarasi konstanta gunakan keyword const
> const ( <br/>
	&nbsp;&nbsp;&nbsp;host   = "localhost"<br/>
	&nbsp;&nbsp;&nbsp;port   = 5432<br/>
	&nbsp;&nbsp;&nbsp;user   = "postgres"<br/>
	&nbsp;&nbsp;&nbsp;pass   = ""<br/>
	&nbsp;&nbsp;&nbsp;dbName = "food"<br/>
)
# to run
> go run main.go

# to compile
> go build


# support
<p><a href="https://www.buymeacoffee.com/rudytoryan"> <img align="left" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="35" width="140" alt="rudytoryan" /></a></p><br><br>