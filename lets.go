package main //Çalıştrılacak her Go kaynak dosyasında yer almalıdır.

// Go Dilinde tek satırlık yorum/açıklama
import (
	"fmt"       // Standart kütüphaneyi projemize dahil ettik
	"io/ioutil" // Temel io işlemleri
	"log"       // log
	"math"      // Matematiksel işlem ve değer tanımlarını içeren kütüphaneyi içe aktardık
	"net/http"
	"os"      // İşletim sistemi işlemleri
	"sort"    // Sıralama işlemleri
	"strconv" // String dönüşümleri
	"strings" // String işlemleri kütüphanesi (Contains, Index, Count vs)
	"time"
)

/*
GO PROGRAMLAMA DİLİ
Go (diğer adıyla golang), Google'da 2007 yılından itibaren geliştirilmeye başlayan açık kaynak programlama dilidir.
Daha çok sistem programlama için tasarlanmış olup, derlenmiş ve statik tipli bir dildir. Kasım 2009'da çıkmıştır. Go
derleyicisi "gc", açık kaynak yazılım olarak, Linux, OS X, Windows, bazı BSD ve Unix versiyonları, ve ayrıca 2015'ten
itibaren akıllı telefonlar için geliştirilmiştir.
Bakınız:
https://youtu.be/rKnDgT73v8s
'Bu bir çok satırlı yorum örneğidir.'
*/

/*
GO PROGRALAMA DİLİNİN SÖZ DİZİMİ
Go'nun sözdizimi, C'den yapılan değişiklikleri içerir, kodu kısa ve okunabilir kılmayı amaçlar. Birleştirilmiş
declaration ve başlatma operatörü, programcılara i := 3 ya da s := "bazı kelimeler" yazarak, herhangi bir tür
belirtimine ihtiyaç duyulmadan değişken tanımlamasını sağlar. Noktalı virgüller hala ifadeleri sonlandırır fakat
satır sonlarında kullanılmasına gerek yoktur. Fonksiyonlar birden fazla değerler döndürebilir (return result, err).
*/

/*
GO DİLİNDE MERHABA DÜNYA PROGRAMI
package main // Çalıştırılacak her proje main paketini içerir
import "fmt" // Standart kütüphaneyi projemize dahil ettik
// Bir fonksiyon tanımı func ifadesi ile başlar onu fonksiyonIsmi() takip eder
// () içerisine parametre tanımı yapılabilir () işartetlerinden sonra eğer varsa
// fonsiyondan dönecek değer türü deklare edilir.
// fonksiyon içersindeki kod bloğu { } işaretleri arasında yer alır.
// Her go programı bir main() fonksiyonu içersinde çalışır
func main() {
	fmt.Println("Merhaba, Dünya!")
}
*/

func main() {
	fmt.Println("Go Eğtimine Hoş Geldiniz") //fmt kütüphanesindeki Println Methodu aldığı stdout'a yazdırır ve bir yani satır ekler
	const pi float64 = 3.14159265           // Açık  deklarasyon ile sabit tanımı
	const egitimTarihi int = 2017
	var yas int = 48 // Açık bir biçimde tür belirterek tanımlama örneği
	var boy float32 = 1.64
	var guzelMi bool = true
	var isim string = "Jennifer Ariston"

	dogumYili := egitimTarihi - yas // Değişkenin taşıdığı değer türü kendiliğinden belirlenecek

	fmt.Println("PI Sayısı:   ", pi) // Bir değişken ile standart satır yazdırma
	// Printf ile verileri biçimlendirip de yazdırırız
	fmt.Printf("Öğrenci adı: %s\nÖğrenci yaşı: %d\nDoğum Yılı: %d\nÖğrenci Boyu: %.2f\nGüzel mi?%t\n", isim, yas, dogumYili, boy, guzelMi)

	// TİPLERİNE GÖRE SAYILAR:
	// uint8  : işaretsiz  8-bit tamsayı (0 ile 255)
	// uint16 : işaretsiz 16-bit tamsayı (0 ile 65535)
	// uint32 : işaretsiz 32-bit tamsayı (0 ile 4294967295)
	// uint64 : işaretsiz 64-bit tamsayı (0 ile 18446744073709551615)
	// int8   : işaretli   8-bit tamsayı (-128 ile 127)
	// int16  : işaretli  16-bit tamsayı (-32768 ile 32767)
	// int32  : işaretli  32-bit tamsayı (-2147483648 ile 2147483647)
	// int64  : işaretli  64-bit tamsayı (-9223372036854775808 ile 9223372036854775807)
	// float32: (-3.4E+38 ile +3.4E+38)
	// float64: (-1.7E+308 ile +1.7E+308) değerleri arasındadır

	// string veri tipi karakter dizilerini içerir
	// bool veri tipi true ve false değerlerini içerir

	fmt.Printf("%d \n", yas)       // int olarak yazımını gösterir
	fmt.Printf("%b \n", dogumYili) // binary olarak yazımını gösterir
	fmt.Printf("%c \n", yas)       // char unicode karakter karşılığını gösterir
	fmt.Printf("%x \n", yas)       //hexedemical
	fmt.Printf("%o \n", yas)       // octal
	fmt.Printf("%e \n", pi)        //scientific yani bilimsel gösterim

	// String İşlemleri
	ornekString := "Merhaba, Dunya!"
	virgulluString := "Metin,Ali,Feyyaz"

	fmt.Println(strings.Contains(ornekString, "Me"))       // String Me içeriyorsa true döner
	fmt.Println(strings.Index(ornekString, "D"))           // D karakterinin endeksi
	fmt.Println(strings.Count(ornekString, "a"))           // String içersindeki boşluklar ve işaretler dahil karakter sayısı
	fmt.Println(strings.Replace(ornekString, "u", "ü", 1)) // u karekteri ü ile değiştirilir
	fmt.Println(strings.Split(virgulluString, ","))        // Stringten dilime çevirme
	// Aşağıda da dilimden stringe çeviriyoruz
	karakterListesi := []string{"z", "y", "x", "w", "v", "u", "t", "s", "r", "q", "p", "o", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
	sort.Strings(karakterListesi)
	fmt.Println("karakterListesi:", karakterListesi)

	// Tür dönüşümleri
	rastgeleBirTamsayi := 3
	rastgeleBirFloat := 3.14
	birBaskaString := "!000000"
	birBaskaStringDaha := "1.37"

	tamSayiydiFloatOldu := float64(rastgeleBirTamsayi)
	floatIkenTamsayiOldu := int(rastgeleBirFloat)

	// Bir string veriyi sayısal değere dönüştürme
	yeniBirTamsayi, _ := strconv.ParseInt(birBaskaString, 0, 64)
	yeniBirFloat, _ := strconv.ParseFloat(birBaskaStringDaha, 64)

	// Go Programlama Dilinin Derleyicisi eğer kullanılmamış bir değişken varsa hata döndürür ve derlemeye izin vermez
	// Gereksiz tanımlamalarla performans kaybı da engellenmiş olur
	// Eğer herhangi bir sebenle kullanılmamış bir tanımlamanız varsa _ ile yakalayabilirsiniz

	_, _, _, _ = tamSayiydiFloatOldu, floatIkenTamsayiOldu, yeniBirTamsayi, yeniBirFloat // Böylece kaynak dosyamız derlenebilecek
	// Elbette değişkenlerin stout ile ekrana yazdırılmaları da değişkeni kullanmak demektir
	// fmt.Print fmt.Printf fmt.Sprintf fmt.Println v.s.

	// Döngüler ve Akış Kontrolleri

	sayac := 0

	fmt.Println("While alternatifi olarak for döngüsü: ")

	for sayac <= 10 {
		fmt.Println(sayac)
		sayac++ // sayac değişkeninin değerini bir arttırır
	}

	fmt.Println("Tipik for döngüsü: ")

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	var yasBilgisi int // Açık deklerasyon ile atama yapmaksızın değişken tanımlama

	fmt.Print("Yaşınızı Giriniz: ")

	fmt.Scanf("%d", &yasBilgisi) //  Kulllanıcı girdisinsini alıp tamsayı formatında değişkene atıyoruz
	// & degiskenin bellek adresine yönlendirir. Pointer örneğinde açıklayacağım

	fmt.Printf("Yaşınız %d olduğuna göre:\n\n", yasBilgisi)

	// If ile akış kontrolu
	if yasBilgisi > 16 && yasBilgisi < 18 {
		fmt.Println("Geçici ehliyet alabilirsiniz ve alkol satılan mekanlara giremezsiniz")
	} else if yasBilgisi >= 18 && yasBilgisi <= 65 {
		fmt.Println("Hem ehliyet alabilir hem alkollü mekanlarda bulanabilirsiniz ama lütfen alkollüyken araç kullanmayın!")
	} else if yasBilgisi > 65 {
		fmt.Println("Ehliyet için her yıl sağlık kontorlünden geçmelisiniz.")
	} else {
		fmt.Println("Hayatın en eğlenceli ve tasasız zamanları :)")
	}

	// swtich/case ile akış kontrolu
	var notBilgisi int

	fmt.Print("Lütfen sınav notunuzu giriniz (0-5): ")
	fmt.Scanf("%d", &notBilgisi)

	fmt.Println("Notunuz:", notBilgisi)

	switch notBilgisi {
	case 0:
		fmt.Println("Kaldın! :(")
	case 1:
		fmt.Println("Pek iyi görünmüyor, çok çalışman gerekecek!")
	case 2:
		fmt.Println("Daha çok çalışmalısın!")
	case 3:
		fmt.Println("Gelecek sefere daha iyisini yaparsın")
	case 4:
		fmt.Println("Başarılı")
	case 5:
		fmt.Println("Çok Başarılı")
	default:
		fmt.Println("Farklı bir not sistemine göre giriş yaptın gibi görünüyor!")
	}

	// Diziler aynı türdeki verilerden oluşan değerlerin gruplardır
	rakamlar := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for iterasyon için de kullanılabilir range ile birlikte
	fmt.Println("for ile iterasyon")

	for _, v := range rakamlar {
		fmt.Println(v)
	}

	diziler := [4]string{"Doctor Who", "American Gods", "Game Of Thrones", "Orange is new black"}

	for i, v := range diziler {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Dilimler dizilerden farklı olarak büyüklükleri değişebilir yani item eklenip çıkarılabilir
	// Diziler ve dilimler 0 ile başlanarak endekslenir

	rakamDilimi := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} // [] dizlerden farklı olarak boş bırakılıyor
	rakamDilimi = append(rakamDilimi, -1, -2, -3)      // dilimimize itemler ekledik

	fmt.Println("rakamDilimi:", rakamDilimi)

	// Go dilinde Ruby ve Python ile kullanılan pop() methodu standat kütüphanede yok
	// Ben de sizin için sil adında aynı işi yapan bir fonksiyon yazdım
	// Tüm fonksiyon, struct, interface tanımlamaları main fonksiyonunun bitiminden sonra yer almaktadır, inceleyebilirsiniz.
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // sondan bir element silen fonksiyonumuzdan dönen değeri atıyoruz
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // böylece dilim bir element kısalıyor
	rakamDilimi = sil(rakamDilimi, len(rakamDilimi)-1) // her seferinde

	fmt.Println("rakamDilimi:", rakamDilimi)

	yarimDilim := rakamDilimi[:5] // Bir endeks aralığını alırken ilk ya da son endeksi belirtmek opsiyoneldir.

	for endeks, deger := range yarimDilim {
		fmt.Printf("Endeks: %d = Değer: %d\n", endeks, deger)
	}

	// make ile dilim oluşturmak
	birDilimDaha := make([]int, 3)

	fmt.Println("birDilimDaha make ile oluşturuldu:", birDilimDaha)

	// Python dilinde dictionary, Ruby dilinde hash olarak adlandırılan anahtar değer eşleşmeleri:
	// yani map veri tipi
	karne := make(map[string]int)
	karne["Matematik"] = 3
	karne["Fizik"] = 4
	karne["Edebiyat"] = 5

	for anahtar, deger := range karne {
		fmt.Printf("%s: %d\n", anahtar, deger)
	}

	karne["Biyoloji"] = 4 // Biyoloji anahtarı 4 tamsayısı ile eşleştirilerek eklendi(pair de diyebiliriz.)

	// Silme işleminden önce
	fmt.Println("Karne:", karne)

	delete(karne, "Biyoloji") // Biyoloji => 4 eşleşmesi map içinden silindi

	// Silme işleminden sonra
	fmt.Println("Karne", karne)

	// main içerisinde fonksiyon çağrısı
	rakamlarToplami := diziDegerleriniTopla(rakamlar) // Görüldüğü gibi bir fonksiyon çağrısın değişkene atayabiliriz

	fmt.Println("rakamlar isimli dizideki değerlerin toplamı:", rakamlarToplami)

	// Go Dilinde fonksiyonlar birden fazla değer dönebilir
	karesi, ikiKati := birdenFazlaDegerDon(4)
	fmt.Printf("Karesi: %d\nİki Katı: %d\n", karesi, ikiKati)

	// Go Dilinde bir fonksia geçeceğimiz argüman saysı belirsizse
	fmt.Println("Birden fazla sayıda argümanın birbiri ile çarpımı", birCokArgumanBirbiriyleCarp(1, 2, 3, 4, 5))

	// Go dilinde closure oluşturma
	birSayiDaha := 2

	katmerle := func() int {
		birSayiDaha *= 2
		return birSayiDaha
	}

	katmerle()
	katmerle()
	katmerle()

	fmt.Println("birSayiDaha:", birSayiDaha)

	// İç içe fonksiyonlar oluşturulabilir
	tamSayi := 7

	fmt.Printf("%d sayısının faktöriyeli %d sayısıdır.\n", tamSayi, faktoriyel(tamSayi))

	// defer ile bir kod grubunun en son çalışması sağlanabilir, hatalar yakalanabilir
	// Son giren ilk çıkar mantığı ile çalıştırılır birden fazla defer
	defer ikiYazdir() // en son çalıştırılacak
	birYazdir()

	// Bir sıfıra bölme hatası yakalayalım
	fmt.Println("Bölme işlemi: 3 / 0 =")
	guvenliBolme(3, 0)
	fmt.Println("Bölme işlemi: 8 / 2 =", guvenliBolme(8, 2))

	// PANIC ile hata yakalama
	panicYakala()

	// değişkenin değerini pointer kullanarak fonksiyon içinde değiştir
	degeriDegisecek := 9

	degeriniKaresiYap(&degeriDegisecek) //& ile bellekteki adresi geçtik

	fmt.Println("degeriDegisecek değişkeninin halihazırdaki değeri:", degeriDegisecek)

	// Kendi veri tiplerimizi struct tanımları içerisinde tasarlayabiliriz
	dGen1 := Dikdortgen{yukseklik: 2.0, en: 3.5}
	dGen1.setYukseklik(4.0)
	dGen1.setEn(7.0)
	cmbr := Cember{yaricap: 2.0}
	cmbr.setYaricap(2.4)
	fmt.Printf("Çember Yarıçap: %.4f\n", cmbr.yaricap)
	fmt.Printf("Dikdörtgen Yükseklik: %.4f\nDikdörtgen En: %.4f\n", dGen1.getYukseklik(), dGen1.getEn())
	fmt.Println("Çemberin Alanı:", alanHesapla(cmbr))
	fmt.Println("Dikdörtgenin Alanı:", alanHesapla(dGen1))

	// Dosya işlemleri
	dosya, err := os.Create("ornek.txt")

	// Dosya oluşturulurken bir hata oluştuysa yakalayıp logladık
	if err != nil {
		log.Fatal(err)
	}

	dosya.WriteString("Dosyamıza string bir veri ekledik.")

	dosya.Close()

	// io/ioutil kütüphanesi ile dosya akışı okuma
	dosyaAkisi, err := ioutil.ReadFile("ornek.txt")

	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}

	dosyayiOku := string(dosyaAkisi)
	fmt.Println("Okundu:", dosyayiOku)

	// Go Routine
	for k := 0; k < 10; k++ {
		go say(k)
	}

	time.Sleep(time.Millisecond * 10000)

	//CHANNEL
	stringCHAN := make(chan string)

	for i := 0; i < 5; i++ {
		go sandvicYap(stringCHAN)
		go mayonezEkle(stringCHAN)
		go hesabaEkle(stringCHAN)
		time.Sleep(time.Second * 2)
	}

	// Web Server
    
    fmt.Println("Web Sunucuyu kapatmak için: CTRL C")

	http.HandleFunc("/yonetici", isleyiciAdmin)
	http.HandleFunc("/", isleyici)

	http.ListenAndServe(":3333", nil) // localhost:3333
}

// INTERFACE, STRUCT, GETTER, SETTER

type Sekil interface {
	alan() float64
}

type Dikdortgen struct {
	yukseklik float64
	en        float64
}

func (dgen Dikdortgen) getYukseklik() float64 {
	return dgen.yukseklik
}

func (dgen *Dikdortgen) setYukseklik(deger float64) {
	dgen.yukseklik = deger
}

func (dgen Dikdortgen) getEn() float64 {
	return dgen.en
}

func (dgen *Dikdortgen) setEn(deger float64) {
	dgen.en = deger
}

func (dgen Dikdortgen) alan() float64 {
	return dgen.yukseklik * dgen.en
}

type Cember struct {
	yaricap float64
}

func (c Cember) getYaricap() float64 {
	return c.yaricap
}

func (c *Cember) setYaricap(deger float64) {
	c.yaricap = deger
}

func (c Cember) alan() float64 {
	return math.Pi * math.Pow(c.yaricap, 2)
}

func alanHesapla(s Sekil) float64 {
	return s.alan()
}

// INTERFACE, STRUCT, GETTER, SETTER TANIMLARI BİTTİ

// FONKSİYON TANIMLARI
func diziDegerleriniTopla(rakamlar [10]int) int {
	toplam := 0

	for _, deger := range rakamlar {
		toplam += deger // toplam = toplam + deger ifadesinin eşleniği
	}
	return toplam
}

func birdenFazlaDegerDon(x int) (int, int) {
	return x * x, x * 2
}

func birCokArgumanBirbiriyleCarp(args ...int) int {
	carpim := 1

	for _, v := range args {
		carpim *= v
	}
	return carpim
}

// Rekürsif (iç içe) fonksiyon
func faktoriyel(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktoriyel(n-1)
}

func birYazdir() {
	fmt.Println("Bir")
}

func ikiYazdir() {
	fmt.Println("İki")
}

func guvenliBolme(sayi1, sayi2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	cozum := sayi1 / sayi2
	return cozum
}

func panicYakala() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("AMAN ALLAHIM, N'OLUYOR!")
}

func degeriniKaresiYap(x *int) {
	*x *= *x // adresteki değere erişip adresteki değerle çarpımını üzerine yazdık
}

func sil(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	return s
}

func say(i int) {
	for j := 0; j < 10; j++ {
		fmt.Println(i, ":", j)
		time.Sleep(time.Millisecond * 1000)
	}
}

// Channel
var sandvicSayisi = 0
var sandvicIsmi = ""

func sandvicYap(stringCHAN chan string) {
	sandvicSayisi++
	sandvicIsmi = "Sandviç No:" + strconv.Itoa(sandvicSayisi)
	fmt.Println(sandvicIsmi, "hazır, mayonez eklemeye gönder!")

	stringCHAN <- sandvicIsmi

	time.Sleep(time.Second * 1)
}

func mayonezEkle(stringCHAN chan string) {
	sandvic := stringCHAN

	fmt.Println("Mayonez eklendi, buyrun:", sandvic)

	stringCHAN <- sandvicIsmi
	time.Sleep(time.Second * 1)
}

func hesabaEkle(stringCHAN chan string) {
	sandvic := <-stringCHAN

	fmt.Println(sandvic, "hesaba eklendi! Müşteriye gönder!")

	time.Sleep(time.Second * 1)
}

// WEB SERVER
func isleyici(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sunucuma Hoş Geldiniz!\n")
}

func isleyiciAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yönetici Sayfasına Hoş Geldiniz(ŞİFRE GEREKTİRİR!)")
}