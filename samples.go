// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


// Lerning materials
// https://www.callicoder.com/golang-structs/
// https://www.calhoun.io/when-nil-isnt-equal-to-nil/  - работа с nil
// Rebuild programm - https://github.com/Gitart/go-watcher
// https://flaviocopes.com/golang-docker/ - DEPLOY TO DOCKER
// https://flaviocopes.com/golang-data-structure-hashtable/   - HASH TABLE

package main

import (
	// "context"
	"fmt"
    // "html/template"
    "text/template"
    "time"
    "os"
    "bytes"
    "crypto/sha1"
    "crypto/md5"
    "crypto/sha256"
    "crypto/hmac"
    "os/user"
    "log"
    "io"
    "strings"
    "math/rand"
    "strconv"
    "encoding/base64"
    "bufio"
    "os/exec"
    "path/filepath"
    	"io/ioutil"
	// "log"
	"net/http"        
    mdl "./model"
    mdd "./model/mdls"

    "github.com/pkg/profile"
)




// ********************************************************************
// Main process
// ********************************************************************
func main() {
     
     defer profile.Start(profile.MemProfile).Stop()

    // Fortune_test()

     Test_randomString()           // RIBUGKVQYN
     randomArray_test()            // [7 2 7 7 6 5 2 3 0 5]
     TestGenRandomArray()          // Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkk=

     // Panic падавление
     // Mpanick()

     // Сортировка когда встречаются слова с большой и маленькой буквы
     // Lists()

    // fmt.Println(join("abc", "123", "xyz"))
	// fmt.Println(joinRunes('a', 'b', 'c', '1', '2', '3', '😁', '😃'))

    // print(Mnth())
    // fmt.Println(StringCode(10))
    // fmt.Println(rand.Int())
    // fmt.Println(rand.Int())

    // fmt.Println(TimeCode())
    // fmt.Println(Md5Strtime("dd","dd"))
    // fmt.Println(M256("Art"))  // 256 hash
     Modeltest_01()     
    // Modeltest_02()
}

// Pusher server
/*
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            if err := pusher.Push("/app.js", nil); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        // ...
    })

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            options := &http.PushOptions{
                Header: http.Header{
                    "Accept-Encoding": r.Header["Accept-Encoding"],
                },
            }
            if err := pusher.Push("/app.js", options); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        // ...
    })

*/

// ***************************************
// Предотвращение паники
// ***************************************
func Mpanick(){
	defer func () {
		if e:=recover(); e!=nil{
			err:=fmt.Errorf("%v",e)
			fmt.Println(err)
		}
	}()

    // Исскуственно создание паники
	panic("bot blin")
}

// ********************************************************************
// Model samples
// Work wit library
// https://www.callicoder.com/golang-methods-tutorial/
// ********************************************************************
func Modeltest_02(){

     rtys:=Mst{"Name":"test"}      // Надо брать в кавычки
     rtys["News"]="News portals"
     fmt.Println(rtys)

     rty:=Mdsa{Name:"test"}
     fmt.Println(rty)


     var Rt mdd.Mdd  

     Rt.Name    = "dddd"
     Rt.Id      = 12456
     Rt.Note    = "Ноте для простого сообщения"
     Rt.Inf.Des = "Описание модели"


     fmt.Println(Rt)

     Rt=mdd.Mdd{122,"Gol","dddd", mdd.Info{"Descrip","Note descr"}}    // НЕ надо брать в кавычки
     fmt.Println(Rt)


     stt:=mdd.Info{"Descrip","Note descr"}
     Rt=mdd.Mdd{122,"Gol","dddd", stt}    // НЕ надо брать в кавычки
     fmt.Println(Rt)

     Rts:=mdd.Mdd{Name: "Вторая замена", Id:5555, Note:"Третьий пример ноте"}
     fmt.Println(Rts)
}	


// ********************************************************************
// Model samples
// ********************************************************************
func Modeltest_01(){
     var r mdl.Utils
     fmt.Println(r.Tn())

     fmt.Println(mdl.Rt.Cds.Itx("Nomer пример"))   // ссылка через перменную опердленную в библиотеке
     fmt.Println(r.Cds.Itx("Nomer пример 333"))    // ссылка локальную перменную
     
     r.Keys.Name="Nomer пример 333"
     d   := r.Keys.Ids("Nomer пример 333")
     i   := r.Keys.Id()
     idn := r.Keys.Mt.Idnm("News IDnm")
     df  := mdl.Rt.Keys.Mt.Idnm("News dff IDnm")

     fmt.Println("IDNM ss : ",df)
     fmt.Println("IDNM : ",idn)
     fmt.Println(d,i)
     fmt.Println(mdl.Rt.Tn())
}


// ********************************************************************
// Model samples
// ********************************************************************
func Modeltest(){
	c := mdl.Customer{Id: 1, Name: "Rajeev Singh", Note:" Переименование переменной на общей"}
	c.Married = true	        // Error: can not refer to unexported field or method
	a := mdl.Address{}          // Error: can not refer to unexported name
	a.HouseNo = "A-123"
	a.City    = "Kiev"

	fmt.Println("Programmer = ", c.Id)
	fmt.Println("Address    = ", a.HouseNo)
}

// ********************************************************************
// Тестирование функций
// ********************************************************************
func Starttest() {
    
    // Случайное число
    Rnds()
    fmt.Println(Int2x(125))
    
    // Хешеривание
    fmt.Println(Md5Codes()) 
    fmt.Println(Md5Str("Art","code"))
    fmt.Println(Md5Str("Art","cods"))
    fmt.Println(Md5Code())
    fmt.Println("md5 file :",Md5file("logs.txt"))
    fmt.Println(Md5file("log.txt"))

    // Sweatrers()	
    // Test_parstxt()
    
    // Запись в файл
    FileWrite("access.log","Open programm..")
    
    GuidStr()
    LookEnv()

    // Переименование файла
    var ct utl
    RenameFileName("access.log", ct.CurTimeStr() + "_newacccess.log")

    // Установка и чтение перменных среды
    Envr()
    fmt.Println(Switching("DOG"))
}


// https://flaviocopes.com/go-random/
// ********************************************************************
// Cлучайное число
// ********************************************************************
func Rnds() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Round ",rand.Intn(1e3))
}

// ********************************************************************
// Тестирование
// ********************************************************************
func Test_parstxt() {

     //  Формирование шаблона  из массива
     gini:= []Mst {{"Country":"Canada","Index":123},{"Country":"Ukrinea","Index":126},{"Country":"France","Index":226} }
     frm := `<table> 
               {{range .}} 
                  {{printf "<tr><td>%s</td><td>%v</td></tr>" .Country .Index}} 
               {{end}} 
             </table>` 
     st  := ParsTxt	(frm,gini)
     fmt.Println(st)

     //  Формирование одной строчки
     ginis:= Mst{"Country":"Canada","Index":123}
     frms := `{{printf "<h1>%s</h1><br><p>%v</p><hr>" .Country .Index}} `
     sts  := ParsTxtSt	(frms,ginis)
     fmt.Println(sts)

     // Crypto
     fmt.Println(Str2Crypto("Любой текст для шифрования в крипто"))
}

// ********************************************************************
// Формирование данных по шаблону и данными
// Данные передаются в виде массива
// Шаблон передается в виде строковой перменной
// На выходе - готовый сформированная строка с данными
// ********************************************************************
func ParsTxt(Parsetxt string, Dat []Mst) string {
	
	buf   := new(bytes.Buffer)
	tmpl,err:= template.New("Maintemplate").Parse(Parsetxt)
	if err!=nil{
	  fmt.Println(err.Error())
	  return "Error read template."
	}
    
    err  = tmpl.ExecuteTemplate(buf, "Maintemplate", Dat)
    if err!=nil{
	  fmt.Println(err.Error())
	  return "Error execute template."
	}
    st  := buf.String()
    return st
}

// ********************************************************************
// Формирование данных по шаблону и данными
// Данные передаются в виде массива
// Шаблон передается в виде строковой перменной
// На выходе - готовый сформированная строка с данными
// ********************************************************************
func ParsTxtSt(Parsetxt string, Dat Mst) string {
	
	buf      := new(bytes.Buffer)
	tmpl,err := template.New("Maintemplate").Parse(Parsetxt)
	if err!=nil{
	  // fmt.Println(err.Error())
	  return "Error read template."
	}

    err  = tmpl.ExecuteTemplate(buf, "Maintemplate", Dat)
    if err!=nil{
	  // fmt.Println(err.Error())
	  return "Error execute template."
	}
    st  := buf.String()
    // buf.Reset() 
    return st
}

// *******************************************************************
// Crypto
// Шифрование в крипто
// fmt.Println(Str2Crypto("Любой текст для шифрования в крипто"))
// *******************************************************************
func Str2Crypto(Txt string) string {
	ab20 := sha1.Sum([]byte(Txt))
	sx16 := fmt.Sprintf("%x", ab20)
	return sx16
}

// *******************************************************************
// Возврат отформатированного темплейта
// *******************************************************************
func Sweatrers() {
	// var ret string
    sweaters  := mdl.Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{ .Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	buf    := new(bytes.Buffer)

	err     = tmpl.ExecuteTemplate(buf, "test", sweaters)
	if err != nil {panic(err) }
	 
	st:=buf.String()
	fmt.Println(st)	
}

// *******************************************************************
// Возврат отформатированного темплейта
// *******************************************************************
func rt() {
	sweaters := mdl.Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{ .Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
	
}

// *******************************************************************
// Возврат отформатированного темплейта
// *******************************************************************
func Pras(){
	gini:=[]mdl.Gini{{"Canada",123},{"Ukrinea",126},{"France",226} }

	gt,err:= template.New("giniTable").Parse(`<table> {{range .}} {{printf "<tr><td>%s</td><td>%v</td></tr>" .Country .Index}} {{end}} </table>`)
	if err!=nil{
		fmt.Println(err.Error())
	}

    gt.Execute(os.Stdout, gini)
    // fmt.Println(ty)
}

// *******************************************************************
// User Info
// *******************************************************************
func GuidStr(){
	usr,_:=user.Current()
	
	f,_:=os.Getwd()
	h,_:=os.Hostname()
	fmt.Println("Uid       :",os.Getuid())
    fmt.Println("Pid       :",os.Getpid())
    fmt.Println("Getwd     :",f)
    fmt.Println("Hostname  :",h)
    fmt.Println("Uid       :",os.TempDir())
	fmt.Println("User Name :",usr.Username)
	fmt.Println("Gid       :",usr.Gid)
	fmt.Println("Uid       :",usr.Uid)
	fmt.Println("Home dir  :",usr.HomeDir)
	fmt.Println("Name      :",usr.Name)
}

// ****************************************************************
// IsExitFile file
// ****************************************************************
func ExitFile(filename string){
	// filename := "a-nonexistent-file"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("file does not exist")
	}
}

// *******************************************************************
// 
// Open and write to file
// If the file doesn't exist, create it, or append to the file
// FileWrite("access.log", "Запись в файл")
// *******************************************************************
func FileWrite(Filename, Appdata  string) {
	

	// f, err := os.OpenFile(Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	f, err := os.OpenFile(Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
	   log.Fatal(err)
	}

    tm:=time.Now().Format("02.01.2006 15:04:05") 
	if _, err := f.Write([]byte(tm+" " + Appdata + "\n")); err != nil {
	  log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// **************************************************************************
// User current
// **************************************************************************
func LookEnv() {
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n",      key, val)
		}
	}

	show("USER")
	show("GOPATH")
}

// **************************************************************************
func RenameFileName(old,new string) {
	os.Rename(old,new)
}

// **************************************************
// Work with date
// **************************************************
type utl struct{}

func (t *utl) Curtime() string {
	return time.Now().Format("02.01.2006 15:04:05") 
}

func (t *utl) Curtimedat() string {
	return time.Now().Format("02012006150405") 
}

func (t *utl) CurTimeStr() string {
	return time.Now().Format("150405") 
}

func (t *utl) OnlyDate() string {
	return time.Now().Format("02.01.2006") 
}

func (t *utl) OnlyTime() string {
	return time.Now().Format("15:04:05") 
}

// **************************************************
// Установка и чтение перменных среды
// **************************************************
func Envr() {
	os.Setenv("TMPDIR", "/my/tmp")
	// defer os.Unsetenv("TMPDIR")
	fmt.Printf("%s lives in %s. MY TEMP DIR %s \n", os.Getenv("USER"), os.Getenv("HOME"),os.Getenv("TEMPDIR"))

}

// **************************************************
// Установка и чтение перменных среды
// **************************************************
func Switching(Code string) string {

 switch Code {
		case "C1","C2","C3":
			return "morning"
		case "USER","DOG" :
			return "Gopher"
}
		return ""
}		

// *****************************************
// Хеш от текущего времени сумма от строковой переменной и соли 
// *****************************************
func TimeCode() string {
	 crutime := time.Now().UnixNano()
	 tm      := strconv.FormatInt(crutime, 10)
	 return tm
}

// *****************************************
// Хеш от текущего времени сумма от строковой переменной и соли 
// *****************************************
func Md5Strtime(name, salt string) string {
	crutime:=time.Now().UnixNano()
    tm:=strconv.FormatInt(crutime, 10)
	h := md5.New()
	io.WriteString(h, tm)
	io.WriteString(h, name)
	io.WriteString(h, salt)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

// *****************************************
// Хеш сумма от строковой переменной и соли 
// *****************************************
func Md5Str(name,salt string) string {
	h := md5.New()
	io.WriteString(h, name)
	io.WriteString(h, salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// *****************************************
// Хеш сумма от времени 
// *****************************************
func Md5Code() string {
    tm:=time.Now().Format("02012006150405")
    sl:=time.Now().Format("150405")
	h := md5.New()
	io.WriteString(h, tm)
	io.WriteString(h, sl)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// *****************************************
// Хеш сумма от времени большими буквами
// *****************************************
func Md5Codes() string {
	cd  := Md5Code()
	ret := strings.ToUpper(cd)
	return ret
}


// *****************************************
// Хеш сумма (Md5) по содержанию файла 
// *****************************************
func Md5file(namefile string) string {
	f, err := os.Open(namefile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

    return fmt.Sprintf("%x", h.Sum(nil))
}

// *****************************************************
// 
// 
// *****************************************************
func M256(t string) string {
	h := sha256.New()
	h.Write([]byte("t\n"))
    return fmt.Sprintf("%x", h.Sum(nil))
}

// *****************************************************
// f:=Cmach2str([]byte("dsdddsss"), []byte("dsddd"))     
// fmt.Println(f)
// *****************************************************
func Cmach2str(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return string(expectedMAC)
}


// *****************************************************
// Cравнение
// *****************************************************
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}


// ************************************************************
// Convert 10->16
// ************************************************************
func Int2x(i int) string {
	// return fmt.Sprintf("0x%04x", i)
	return fmt.Sprintf("0x%x", i) + "-----" +fmt.Sprintf("0x%8x", i)
}


// ************************************************************
// 10 -> 8 
// ************************************************************
func Int8x(i int) string {
	return fmt.Sprintf("0x%8x", i)
}


// ************************************************************
// Функция для восстановления сбойной ситуации
// ************************************************************
/*
func Namefuncs() (i int, errr error) {
	defer func () {
		if e:=recover(); e!=nil{
		   err=fmt.Errorf("%v",e)
		   fmt.Println(err)
		}
	}()
	i=Convert(x)
	return i,nil
}
*/

// ************************************************************
// Функция для восстановления сбойной ситуации
// ************************************************************
/*
func LogPage(function func(http.ResponseWrite, *http.Requset) ) func (http.ResponseWrite, *http.Requset) {
	return func (Writer http.ResponseWrite, Requset *http.Requset){
		defer func () {
		      if x:=recover(); x!=nil{
		      	 log.Printf("[%v] panic: %v ", Requset.RemoteAddr(), x)
		      }	
		}()
		function (Writer, Requset)
    }
}
*/


// ****************************************************************************
// Формирование уникального кода на основании символов которые подаются на вход
// https://www.calhoun.io/creating-random-strings-in-go/
// ****************************************************************************
func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

// ************************************************
// Test StringCode
// ************************************************
func StringCode(length int) string {
     // Используемые сиволы в коде
     var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
     // var charset = "dlASDFGXYZ0123456789"
     rets:= StringWithCharset(length, charset)
     rets=strings.ToUpper(rets)
     return rets
}

// ************************************************
// 
// Опредление кода на соновании даты
// https://medium.com/@kpbird/golang-generate-fixed-size-random-string-dd6dbd5e63c0
// https://play.golang.org/p/j_c4BDoIUw
// ************************************************
func Mnth() string {
	m:=time.Now().Format("1")
    d:=time.Now().Format("2") 
     s,_:=strconv.Atoi(d)
     f:=s+1

    a:=map[string]string{"1":"dd","3":"ddd","4":"четв"}
    var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    fmt.Println("Текущему дню соответствует буква из массива ",charset[s:f], s)

	return a[d]+m
}

// ************************************************
// Test Random Code
// ************************************************
func RandomString_Test() {
    fmt.Println("Random String Example")
    fmt.Println()
    fmt.Println("20 chars: " + RandomString(20))
    fmt.Println("10 chars: " + RandomString(10))
    fmt.Println("90 chars: " + RandomString(90))
}

// ************************************************
// Random Code
// ************************************************
func RandomString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}

// ************************************************
// Использование буилдора начиная с 1.10 версии
// https://www.calhoun.io/concatenating-and-building-strings-in-go/
// ************************************************
func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func joinRunes(runes ...rune) string {
	var sb strings.Builder
	for _, r := range runes {
		sb.WriteRune(r)
	}
	return sb.String()
}

// CRYPTO-LEVEL RANDOM NUMBERS
func TestGenRandomArray(){

// Example: this will give us a 44 byte, base64 encoded output
token, err := GenerateRandomString(32)
if err != nil {
	// Serve an appropriately vague error to the
    // user, but log the details internally.
}
fmt.Println(token)
}

// *******************************************************************************************
// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// GenerateRandomBytes returns securely generated random bytes. 
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// 
// Output :
// cmLfS5xc9VRjaaw7-XTSAoRvDO-RRFGR4D3Zxfmbnak=
// *******************************************************************************************
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// *******************************************************************************************
// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// *******************************************************************************************
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}


// *******************************************************************************************
// GENERATE A RANDOM ARRAY OF INTEGERS
// [3 9 7 4 4 8 2 4 6 0]
// *******************************************************************************************
func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}

func randomArray_test() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(randomArray(10))
}

// ********************************************************************
// GENERATE A RANDOM STRING
// ********************************************************************
// https://flaviocopes.com/go-random/
// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}


func Test_randomString() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(randomString(10)) // print 10 chars
}

// ****************************************************************
// Вычисление количество месяцев между датой и текущей датой
// monthsCountSince calculates the months between now
// and the createdAtTime time.Time value passed
//...
// createdAt := time.Date(2011, 9, 2, 0, 0, 0, 0, time.UTC)
// count = monthsCountSince(createdAt)
// ****************************************************************
func monthsCountSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

// ****************************************************************
// HANDLING PRE-FLIGHT OPTIONS REQUESTS
// CORS does pre-flight requests sending an OPTIONS request to any URL, so to handle a POST request you first need to handle an OPTIONS request to that same URL.
// On GET endpoints this is not an issue, but it is for POST and PUT and DELETE requests.
// How?
// https://flaviocopes.com/golang-enable-cors/
// ****************************************************************
/*
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

    // process the request...
}
*/

// ****************************************************************
// remove `3` in 2 different ways. output:
// [1 2 4 5 6 7]
// [1 2 4 5 6 7]
// ****************************************************************
func AppendTest() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 2, 3, 4, 5, 6, 7}
	i := 2
	a = append(a[:i], a[i+1], a[i+2], a[i+3], a[i+4])
	b = append(b[:i], b[i+1:]...)

	fmt.Println(a)
	fmt.Println(b)
}

type Animal struct {
	Age interface{}
}

// "3" string
// 3 int
// "not really an age" string
func StructureTest() {
	dog := Animal{}
	dog.Age = "3"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = 3
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = "not really an age"
	fmt.Printf("%#v %T", dog.Age, dog.Age)
}

var files []string

func visit(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, "/off/") {
		return nil
	}
	if filepath.Ext(path) == ".dat" {
		return nil
	}
	if f.IsDir() {
		return nil
	}
	files = append(files, path)
	return nil
}

func Fortune_test() {
	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}
	fortuneCommand.Start()
	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()
	line := outputStream.Text()
	root := line[strings.Index(line, "/"):]

	err = filepath.Walk(root, visit)
	if err != nil {
		panic(err)
	}
}

// *****************************************************************
// List files
// https://flaviocopes.com/go-list-files/
// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string) ([]os.FileInfo, error) {
    f, err := os.Open(dirname)
    if err != nil {
        return nil, err
    }
    list, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        return nil, err
    }
    // sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
    return list, nil
}

func visits(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        *files = append(*files, path)
        return nil
    }
}

func TestReadFiles() {
    var files []string

    root := "/some/folder/to/scan"
    err := filepath.Walk(root, visits(&files))
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}

// 
func readbodyGet() {
	res, err := http.Get("https://api.github.com/search/repositories?q=stars:>=10000+language:go&sort=stars&order=desc")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	log.Printf("Body: %s\n", body)
}

//The following snippet shows how to print the content of a slice of strings into a file, separating each string with a new line.

// func Redfiles() {
//     var text []string

//     //... fill text with string values

//     // writes a slice of strings to a file, one per line
//     sep := "\n"
//     for _, line := range text {
//         if _, err = f.WriteString(line + sep); err != nil {
//             panic(err)
//         }
//     }
// }


// calcOffset determines and returns the amount of days missing to fill
// the last row of the stats graph
func calcOffset() int {
    var offset int
    weekday := time.Now().Weekday()

    switch weekday {
    case time.Sunday:    offset = 7
    case time.Monday:    offset = 6
    case time.Tuesday:   offset = 5
    case time.Wednesday: offset = 4
    case time.Thursday:  offset = 3
    case time.Friday:    offset = 2
    case time.Saturday:  offset = 1
    }
    return offset
}

// // SORT THE MAP
// // sortMapIntoSlice returns a slice of indexes of a map, ordered
// func sortMapIntoSlice(m map[int]int) []int {
//     // order map
//     // To store the keys in slice in sorted order
//     var keys []int
//     for k := range m {
//         keys = append(keys, k)
//     }
//     sort.Ints(keys)

//     return keys
// }
// sortMapIntoSlice() takes a map and returns a slice with the map keys ordered by their integer value. This is used to print the map properly sorted.

// // GENERATE THE COLUMNS
// // buildCols generates a map with rows and columns ready to be printed to screen
// func buildCols(keys []int, commits map[int]int) map[int]column {
//     cols := make(map[int]column)
//     col := column{}

//     for _, k := range keys {
//         week := int(k / 7) //26,25...1
//         dayinweek := k % 7 // 0,1,2,3,4,5,6

//         if dayinweek == 0 { //reset
//             col = column{}
//         }

//         col = append(col, commits[k])

//         if dayinweek == 6 {
//             cols[week] = col
//         }
//     }

//     return cols
// }
