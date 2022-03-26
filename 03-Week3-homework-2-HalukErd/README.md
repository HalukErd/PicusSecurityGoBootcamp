# Medium Article About Concurrency in Go
https://medium.com/@erd.haluk/concurrency-in-golang-how-and-how-come-is-different-e702a59156a0

# Explanation
## directories
- I have created couple packages to separate concerns.
```text
.
├── README.md
├── books
├── command
│   └── command.go
├── core
│   └── book.go
├── go.mod
├── main.go
├── repository
│   └── bookRepository.go
└── service
    └── bookService.go
```
## paramParser
- I have added paramParser logic to determine how to parse params for specific command.
```go
var paramParser = make(map[string]func() Params)

func PopulateParamParser() {
	paramParser["search"] = ReadNameParam
	paramParser["get"] = ReadBookIdParam
	paramParser["delete"] = ReadBookIdParam
	paramParser["buy"] = ReadIdAndQtyParam
}

func HandleReadParams(c Cmd) Params {
	paramParserFunc, ok := paramParser[c.Key]
	if !ok {
		return Params{}
	}
	return paramParserFunc()
}
```

## types for entities
- I have added a BaseEntity because every entity should have some fundamental attributes like id, name, creationTime, modifiedTime etc.
```go
type BaseEntity struct {
	Id   int64
	Name string
}

type Author struct {
	BaseEntity
}

type Book struct {
	BaseEntity
	PageNumber int
	Stock      int
	Price      float64
	StockCode  int
	Isbn       string
	Author     Author
}
```



# Test Results
## Wrong Usage Test Results
- go run main.go WRONG-COMMAND-NAME
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go add newBook
Could not run 'add' command.
Available Commands are below.
list
search
get
delete
buy
```

- go run main.go get noBookId
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go get        
The method needs bookID param.
```

## Happy Path Test Results
- go run main.go LIST
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go list
Id:0  Name:A Song Of Ice And Fire  Page Number:1200  Stock:253  Price:23.500000  Stock Code:1513  ISBN:99921-58-10-7  Author:George RR Martin
Id:1  Name:The Kingkiller Chronicle  Page Number:623  Stock:123  Price:54.200000  Stock Code:12314  ISBN:9971-5-0210-0  Author:Patrick Rothfus
Id:2  Name:The Lord Of The Rings  Page Number:1366  Stock:121  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien
```

- go run main.go SEARCH bookName
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go search Lord of the Ring                    
---Search Result---
Id:2  Name:The Lord Of The Rings  Page Number:1366  Stock:121  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go search American Gods   
Your search: 'American Gods' could not be found


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go search of
---Search Result---
Id:0  Name:A Song Of Ice And Fire  Page Number:1200  Stock:253  Price:23.500000  Stock Code:1513  ISBN:99921-58-10-7  Author:George RR Martin
Id:2  Name:The Lord Of The Rings  Page Number:1366  Stock:121  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien
```

- go run main.go DELETE bookId
````text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go delete 2
Book is deleted.
Updated Book List
Id:0  Name:A Song Of Ice And Fire  Page Number:1200  Stock:253  Price:23.500000  Stock Code:1513  ISBN:99921-58-10-7  Author:George RR Martin
Id:1  Name:The Kingkiller Chronicle  Page Number:623  Stock:123  Price:54.200000  Stock Code:12314  ISBN:9971-5-0210-0  Author:Patrick Rothfus


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go delete 3
Could not delete book with id:3
````

- go run main.go GET bookId
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go get 2
Id:2  Name:The Lord Of The Rings  Page Number:1366  Stock:121  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go get 4
Could not found
```

- go run main.go BUY bookId bookQty
```text
➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go list
Id:0  Name:A Song Of Ice And Fire  Page Number:1200  Stock:253  Price:23.500000  Stock Code:1513  ISBN:99921-58-10-7  Author:George RR Martin
Id:1  Name:The Kingkiller Chronicle  Page Number:623  Stock:123  Price:54.200000  Stock Code:12314  ISBN:9971-5-0210-0  Author:Patrick Rothfus
Id:2  Name:The Lord Of The Rings  Page Number:1366  >>> Stock:121 <<<  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go buy 2 50
Buy Book process has finished with success.
Updated Book List
Id:0  Name:A Song Of Ice And Fire  Page Number:1200  Stock:253  Price:23.500000  Stock Code:1513  ISBN:99921-58-10-7  Author:George RR Martin
Id:1  Name:The Kingkiller Chronicle  Page Number:623  Stock:123  Price:54.200000  Stock Code:12314  ISBN:9971-5-0210-0  Author:Patrick Rothfus
Id:2  Name:The Lord Of The Rings  Page Number:1366  >>> Stock:71 <<<  Price:52.800000  Stock Code:2131  ISBN:960-425-059-0  Author:Tolkien


➜  homework-2-week-3-HalukErd git:(main) ✗ go run main.go buy 2 150
There is low stock.
```

## Homework | Week 3
`Not: Ödevi yeni bir repoya ekleyeceksiniz. Var olan reponuzda bir güncelleme olmayacak. "homework-2..." şeklinde yeni bir repo üzerinde çalışacaksınız.`

Elimizde bir kitap listesi var. 
Kitap alanları şöyle;
```
- Kitap ID
- Kitap Adı
- Sayfa Sayısı
- Stok Sayısı
- Fiyatı
- Stock Kodu
- ISBN
- Yazar bilgisi (ID ve İsim)
```

1. Tüm kitapları listele (list)
2. Verilen girdi hangi kitap isimlerinde geçiyorsa o kitapları listele (search)
3. ID'ye göre kitabı yazdır
4. IDsi verilen kitabı sil. (Silinen kitabın ID'ye göre geliyor olması gerekiyor.)
5. IDsi verilen kitabı istenilen adet kadar satın al ve kitabın son bilgilerini ekrana yazdır.

Yanlış komut girildiğinde ekrana usage'ı yazdıracak. 


Concurrency ile ilgili medium yazısı yazılacak. 

### list command
```
go run main.go list
```

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of the King
```

### get command
```
go run main.go get <bookID>
go run main.go get 5
```

### delete command
```
go run main.go delete <bookID>
go run main.go delete 5
```

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```

###
# Requirements:
- README
- No third party package(s)
- Everything should be in English (Comments, Function names, File names, etc.)
- Use structs not maps

# Test Results
