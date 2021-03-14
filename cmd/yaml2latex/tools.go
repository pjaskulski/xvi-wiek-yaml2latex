package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

var monthName = map[string]string{
	"01": "Styczeń",
	"02": "Luty",
	"03": "Marzec",
	"04": "Kwiecień",
	"05": "Maj",
	"06": "Czerwiec",
	"07": "Lipiec",
	"08": "Sierpień",
	"09": "Wrzesień",
	"10": "Październik",
	"11": "Listopad",
	"12": "Grudzień",
}

// readFact func
func (app *application) readFact(filename string) {
	var result []Fact
	var fact Fact

	name := filenameWithoutExtension(filepath.Base(filename))

	fileBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	r := bytes.NewReader(fileBuf)
	yamlDec := yaml.NewDecoder(r)

	yamlErr := yamlDec.Decode(&fact)

	for yamlErr == nil {
		/* walidacja danych w strukturze fact (część pól jest wymaganych, brak nie
		   zatrzymuje działania aplikacji, ale jest odnotowywany w logu).
		*/
		err = fact.Validate()
		if err != nil {
			app.errorLog.Println("file:", filepath.Base(filename)+",", "error:", err)
		}

		fact.ContentLatex = prepareFactLatex(fact.Content, fact.ID, fact.Sources)

		if fact.Geo != "" {
			fact.GeoLatex = prepareGeoLatex(fact.Geo, fact.Location)
		} else if fact.Location != "" {
			fact.GeoLatex = `\vspace{2mm} 
\noindent Miejsce wydarzenia: ` + fact.Location
		}

		result = append(result, fact)

		yamlErr = yamlDec.Decode(&fact)
	}

	// jeżeli był błąd w pliku yaml, inny niż koniec pliku to zapis w logu
	if yamlErr != nil && yamlErr.Error() != "EOF" {
		app.errorLog.Println("file:", filepath.Base(filename)+",", "error:", yamlErr)
	}

	numberOfFacts += len(result)
	app.dataCache[name] = result
}

// loadData - wczytuje podczas startu serwera dane do struktur w pamięci operacyjnej
func (app *application) loadData(path string) error {
	// wydarzenia
	app.infoLog.Printf("Wczytywanie bazy wydarzeń...")
	start := time.Now()

	dataFiles, _ := filepath.Glob(filepath.Join(path, "*-*.yaml"))
	if len(dataFiles) > 0 {
		for _, tFile := range dataFiles {
			app.readFact(tFile)
		}
	}

	elapsed := time.Since(start)
	app.infoLog.Printf("Czas wczytywania danych: %s", elapsed)

	return nil
}

/* dayFact - funkcja zwraca fragment html z linkiem jeżeli dla danego dnia są wydarzenia
   historyczne w bazie, lub sam numer dnia (o szarym kolorze) jeżeli ich nie ma.
   Wykorzystywana w kalendarzu.
*/

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func prepareGeoLatex(geo string, location string) string {
	textLatex := `\vspace{2mm} 
\noindent Miejsce wydarzenia: \href{[geo]}{[location]}`
	pos := strings.Split(geo, ",")
	if len(pos) == 2 {
		url := strings.Replace(`https://www.openstreetmap.org/?mlat=[lat]&mlon=[lon]&zoom=9`, "[lat]", pos[0], 1)
		url = strings.Replace(url, "[lon]", pos[1], 1)
		textLatex = strings.Replace(textLatex, "[geo]", url, 1)
		textLatex = strings.Replace(textLatex, "[location]", location, 1)
	}
	return textLatex
}

func prepareTextStyle(content string, clear bool) string {
	// kapitaliki
	var rgx = regexp.MustCompile(`\{\{\{(.*?)\}\}\}`)
	var pre, post string
	if !clear {
		pre = `<newthought>`
		post = `</newthought>`
	}

	textToSmallCaps := rgx.FindAllString(content, -1)

	for _, item := range textToSmallCaps {
		textLatex := pre + item[3:len(item)-3] + post
		content = strings.Replace(content, item, textLatex, -1)
	}

	// pogrubienie
	var rgxb = regexp.MustCompile(`\{\{(.*?)\}\}`)
	if !clear {
		pre = `<textbf>`
		post = `</textbf>`
	}

	textBold := rgxb.FindAllString(content, -1)

	for _, item := range textBold {
		textLatex := pre + item[2:len(item)-2] + post
		content = strings.Replace(content, item, textLatex, -1)
	}

	// italiki
	var rgxi = regexp.MustCompile(`\{(.*?)\}`)
	if !clear {
		pre = `\emph{`
		post = `}`
	}

	textItalic := rgxi.FindAllString(content, -1)

	for _, item := range textItalic {
		textLatex := pre + item[1:len(item)-1] + post
		content = strings.Replace(content, item, textLatex, -1)
	}

	content = strings.Replace(content, `<newthought>`, `\newthought{`, -1)
	content = strings.Replace(content, `</newthought>`, `}`, -1)
	content = strings.Replace(content, `<textbf>`, `\textbf{`, -1)
	content = strings.Replace(content, `</textbf>`, `}`, -1)

	return content
}

func prepareFactLatex(content string, id string, sources []Source) string {

	content = prepareTextStyle(content, false)

	pre := `\sidenote{`
	post := `}`

	for _, item := range sources {
		value := prepareTextStyle(item.Value, false)
		if item.URL != "" {
			var nameURL string
			if item.URLName != "" {
				nameURL = item.URLName
			} else {
				nameURL = item.URL
			}
			value += fmt.Sprintf(` \href{%s}{%s}`, item.URL, nameURL)
			value = strings.Replace(value, `#`, `\#`, -1)
			value = strings.Replace(value, `%`, `\%`, -1)
			value = strings.Replace(value, `_`, `\_`, -1)
		}
		content = strings.Replace(content, "["+item.ID+"]", pre+value+post, -1)
	}

	return content
}

// isRunByRun - funkcja sprawdza czy uruchomiono program przez go run
// czy też program skompilowany, funkcja dla systemu Linux
func isRunByRun() bool {
	if strings.Index(os.Args[0], "/tmp/go-build") != -1 {
		return true
	}
	return false
}

func (app *application) writeLatex(text string) {
	_, err := app.f.WriteString(text + "\n")
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

func (app *application) createLatex(filename string) {
	var err error

	app.f, err = os.Create(filename)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	defer app.f.Close()

	// początek
	app.writeLatex(latex_begin)

	// treść
	currentMonth := "01"
	app.writeLatex(fmt.Sprintf(`\chapter{%s}`, monthName[currentMonth]))
	app.writeLatex("")

	quote := quotes[currentMonth]
	app.writeLatex(fmt.Sprintf(latex_quote, quote.Content, quote.Source))

	var factKeys []string
	for key := range app.dataCache {
		factKeys = append(factKeys, key)
	}
	sort.Strings(factKeys)

	for _, item := range factKeys {
		if item[:2] != currentMonth {
			app.writeLatex(fmt.Sprintf(`\chapter{%s}`, monthName[item[:2]]))
			currentMonth = item[:2]
			quote := quotes[currentMonth]
			app.writeLatex(fmt.Sprintf(latex_quote, quote.Content, quote.Source))
		}
		facts := app.dataCache[item]

		for _, fact := range facts {
			factDate := fmt.Sprintf("%d %s %d", fact.Day, strings.ToLower(monthName[item[:2]]), fact.Year)
			app.writeLatex(fmt.Sprintf(`\section[%s]{%s}`, factDate+" - "+fact.Title, fact.Title))
			app.writeLatex(fact.ContentLatex)
			if fact.People != "" {
				persons := strings.Split(fact.People, ";")
				for _, person := range persons {
					person = strings.TrimSpace(person)
					app.writeLatex(fmt.Sprintf(`\index[person]{%s}`, person))
				}
			}
			if fact.GeoLatex != "" {
				app.writeLatex("")
				app.writeLatex(fact.GeoLatex)
			}
			app.writeLatex("")
			app.writeLatex(`\noindent\hrulefill`)
			app.writeLatex(`\vspace{2mm}`)
			app.writeLatex("")
		}
	}

	// zakończenie
	app.writeLatex(latex_end)
}
