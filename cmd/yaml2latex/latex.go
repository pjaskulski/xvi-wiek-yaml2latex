package main

type Quote struct {
	Content string
	Source  string
}

var latex_begin = `
\documentclass[a4paper]{tufte-book}

\hypersetup{colorlinks}
\widowpenalty=1000
\clubpenalty=1000

\usepackage{polski}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage{tgpagella}
\usepackage{cmap}
\usepackage{imakeidx}

\title{XVI wiek\\ \noindent{Kalendarium}}
\author[Piotr Jaskulski]{Opracowanie: Piotr Jaskulski}
\publisher{Serwis xvi-wiek.pl}

\usepackage{fancyvrb}
\fvset{fontsize=\normalsize}

\newcommand{\hangp}[1]{\makebox[0pt][r]{(}#1\makebox[0pt][l]{)}}
\newcommand{\hangstar}{\makebox[0pt][l]{*}}

\usepackage{xspace}

\newcommand{\monthyear}{%
  \ifcase\month\or styczniu\or lutym\or marcu\or kwietniu\or maju\or czerwcu\or
  lipcu\or sierpniu\or wrześniu\or październiku\or listopadzie\or
  grudniu\fi\space\number\year
}

\newcommand{\openepigraph}[2]{%
  \begin{fullwidth}
  \sffamily\large
  \begin{doublespace}
  \noindent\allcaps{#1}\\
  \noindent\allcaps{#2}
  \end{doublespace}
  \end{fullwidth}
}

\newcommand{\blankpage}{\newpage\hbox{}\thispagestyle{empty}\newpage}

\usepackage{units}

\newcommand{\measure}[3]{#1/#2$\times$\unit[#3]{pc}}

\newcommand{\hlred}[1]{\textcolor{Maroon}{#1}}
\newcommand{\hangleft}[1]{\makebox[0pt][r]{#1}}
\newcommand{\hairsp}{\hspace{1pt}}
\newcommand{\hquad}{\hskip0.5em\relax}
\newcommand{\TODO}{\textcolor{red}{\bf TODO!}\xspace}
\newcommand{\na}{\quad--}
\providecommand{\XeLaTeX}{X\lower.5ex\hbox{\kern-0.15em\reflectbox{E}}\kern-0.1em\LaTeX}
\newcommand{\tXeLaTeX}{\XeLaTeX\index{XeLaTeX@\protect\XeLaTeX}}
\newcommand{\tuftebs}{\symbol{'134}}
\newcommand{\doccmdnoindex}[2][]{\texttt{\tuftebs#2}}
\newcommand{\doccmddef}[2][]{%
  \hlred{\texttt{\tuftebs#2}}\label{cmd:#2}%
  \ifthenelse{\isempty{#1}}%
    {
      \index{#2 command@\protect\hangleft{\texttt{\tuftebs}}\texttt{#2}}
    }
    {
      \index{#2 command@\protect\hangleft{\texttt{\tuftebs}}\texttt{#2} (\texttt{#1} package)}
      \index{#1 package@\texttt{#1} package}\index{packages!#1@\texttt{#1}}
    }
}
\newcommand{\doccmd}[2][]{%
  \texttt{\tuftebs#2}%
  \ifthenelse{\isempty{#1}}%
    {
      \index{#2 command@\protect\hangleft{\texttt{\tuftebs}}\texttt{#2}}
    }
    {
      \index{#2 command@\protect\hangleft{\texttt{\tuftebs}}\texttt{#2} (\texttt{#1} package)}
      \index{#1 package@\texttt{#1} package}\index{packages!#1@\texttt{#1}}
    }%
}
\newcommand{\docopt}[1]{\ensuremath{\langle}\textrm{\textit{#1}}\ensuremath{\rangle}}%
\newcommand{\docarg}[1]{\textrm{\textit{#1}}}
\newenvironment{docspec}{\begin{quotation}\ttfamily\parskip0pt\parindent0pt\ignorespaces}{\end{quotation}}
\newcommand{\docenv}[1]{\texttt{#1}\index{#1 environment@\texttt{#1} environment}\index{environments!#1@\texttt{#1}}}
\newcommand{\docenvdef}[1]{\hlred{\texttt{#1}}\label{env:#1}\index{#1 environment@\texttt{#1} environment}\index{environments!#1@\texttt{#1}}}
\newcommand{\docpkg}[1]{\texttt{#1}\index{#1 package@\texttt{#1} package}\index{packages!#1@\texttt{#1}}}
\newcommand{\doccls}[1]{\texttt{#1}}
\newcommand{\docclsopt}[1]{\texttt{#1}\index{#1 class option@\texttt{#1} class option}\index{class options!#1@\texttt{#1}}}
\newcommand{\docclsoptdef}[1]{\hlred{\texttt{#1}}\label{clsopt:#1}\index{#1 class option@\texttt{#1} class option}\index{class options!#1@\texttt{#1}}}
\newcommand{\docmsg}[2]{\bigskip\begin{fullwidth}\noindent\ttfamily#1\end{fullwidth}\medskip\par\noindent#2}
\newcommand{\docfilehook}[2]{\texttt{#1}\index{file hooks!#2}\index{#1@\texttt{#1}}}
\newcommand{\doccounter}[1]{\texttt{#1}\index{#1 counter@\texttt{#1} counter}}

\makeindex[name=person,title={Indeks postaci}]

\begin{document}

\frontmatter

\maketitle

\newpage
\begin{fullwidth}
~\vfill
\thispagestyle{empty}
\setlength{\parindent}{0pt}
\setlength{\parskip}{\baselineskip}

\par\smallcaps{Opublikowane przez serwis xvi-wiek.pl}

\par Zastrzeżenie: autor dołożył wszelkich starań, by zawarte w tej książce informacje 
były kompletne i rzetelne. Nie bierze jednak żadnej odpowiedzialności za ich wykorzystanie, 
nie ponosi również żadnej odpowiedzialności za ewentualne szkody wynikłe z wykorzystania 
informacji zawartych w treściach opublikowanych w niniejszej publikacji.

\par Opisy wydarzeń są dostępne na licencji Creative Commons Uznanie 
Autorstwa 4.0 Międzynarodowe.
\par\textit{Plik wygenerowany w \monthyear}
\end{fullwidth}

% dedykacja
\cleardoublepage
~\vfill
\thispagestyle{empty}
\begin{doublespace}
\noindent\fontsize{27}{33}\selectfont\itshape
\nohyphenation
Dominice
\end{doublespace}
\vfill
\vfill

\cleardoublepage

\mainmatter

\chapter{Przedmowa}
\begin{fullwidth}
{\large Przedstawione Czytelnikowi opracowanie (kalendarium) zawiera kolekcję informacji 
dotyczących wydarzeń z okresu zbliżonego zakresem do XVI wieku, potocznie nazywanego 
złotym wiekiem dziejów Polski. Jednakże dokładny zakres chronologiczny obejmuje 
lata 1490-1586, rok 1490 to rok urodzin Albrechta Hohenzollerna, 1586 z kolei 
to rok śmierci Stefana Batorego. Nie wynika to z żadnych innych przesłanek niż 
zainteresowania autora. Geograficznie prezentowane fakty i ciekawostki ograniczają 
się głównie do zasięgu ówczesnego Królestwa Polskiego i Prus Książęcych, lecz 
także krain sąsiednich jeżeli w wydarzenia na ich terenie Królestwo Polskie było 
zaangażowane lub dane wydarzenie na losy Królestwa miało wpływ.

Technicznie ebook jest przetransferowaną do formy książki zawartością serwisu 
xvi-wiek.pl. Przyczyny powstania tego ebooka i serwisu xvi-wiek.pl były tak naprawdę 
dwie: XVI wiek to mój ulubiony okres historii Polski a programowanie aplikacji 
webowych w języku Go od pewnego czasu zajmowało pierwsze miejsce na liście technologii 
do nauczenia. Kiedy moja córka otrzymała ocenę dostateczną z klasówki z historii, 
właśnie z XVI wieku, poczułem się wywołany do tablicy, powstała strona www oraz
mechanizm automatycznej konwersji ze źródeł strony do tej książki...}
\end{fullwidth}
`

var latex_end = `
\backmatter

\printindex[person] 

\end{document}
`

var latex_quote = `
\begin{fullwidth}
\emph{%s} 
\vspace{5mm}
\begin{flushright} 
\small{%s}
\end{flushright} 
\end{fullwidth}
\vspace{20mm}
`

var quotes = map[string]Quote{
	"01": {
		`Rzeczpospolita niczym innym w całości i dłużej zachowana być nie może, 
jedno zgodą, miłością, społecznością, jednością, bo jak mądrze ktoś [...] napisał 
'zgodą małe rzeczy rosną, niezgodą wielkie upadają' i 'moc zjednoczenia mocniejsza 
niż rozdwojenia'[...]`,
		`fragment testamentu Zygmunta II Augusta, za: U. Borkowska, "Dynastia Jagiellonów w Polsce", 2011`,
	},
	"02": {
		`Gdy taki jest los i przeznaczenie rzeczy ludzkich, że zgoda i pokój między 
ludźmi rozmaitych królestw i prowincyj wiecznie trwać nie mogą, i że po 
długotrwałym pokoju, wojna niekiedy następować musi, a sama zmienność 
rzeczy uczy, iż każde królestwo, nie mniej obawiać się drugiej, jak pierwszego
spodziewać się powinno, najlepszą tedy, a rządcom królestw i prowincyj 
arcypożyteczną rzeczą jest, aby się zawczasu starali o to, iżby kraje i ludy im 
poddane, na wojnę , którąby im mógł wydać nieprzyjaciel zawsze były przygotowane,
i do wszelakiego odporu przysposobione. `,
		`Jan Tarnowski, "Consilium rationis bellicae", 1558, Kraków, (wydanie z 1858 r.) str. 5`,
	},
	"03": {
		`Brak wyobraźni, zwłaszcza w materii politycznej, wiedzie do gorszych rzeczy
niż zbrodnia - do błędów`,
		`Grzegorz Kucharczyk, "Prusy. Pięć wieków", 2020`,
	},
	"04": {
		`Od niepamiętnych czasów istnieją w polityce pewne zasady, ustalone bez
teoretyzowania, na mocy samej praktyki, wielokrotnie sprawdzone doświadczeniem.
Przestrzeganie ich świadczy o kulturze politycznej kraju. Głoszą one, iż 
nie wolno wojować na wielu frontach jednocześnie, że trwałe korzyści osiąga 
się idąc naprzód krok za krokiem - "etapami", jak zaczęto mówić w czasach
najnowszych - że nie sztuka chwycić rzecz, która chwilowo ma słabego dozorcę, 
sztuka utrzymać... `,
		`Paweł Jasienica, "Polska Jagiellonów", 1965, str. 266`,
	},
	"05": {
		`Królestwo Polskie nadal było jednak na tle krajów Zachodu krajem słabo 
zurbanizowanym. Miasta budowano z drewna, nie z kamienia. W 1400 r. w 
Królestwie Polskim było około 700 murowanych budowli (w całym!), w 1500 r.
- zaledwie około 1350.`,
		`Adam Leszczyński, "Ludowa historia Polski", 2020, str. 99`,
	},
	"06": {
		`Teraz porządki francuskie chciał Henryk [Walezy] zaprowadzić pod Wawelem. Rychło 
przekonał się o tym prymas Uchański, gdy chciał postawić przed sądem duchownym 
jednego z księży apostatów i uprowadził go siłą wbrew straży marszałkowskiej. 
Na rozkaz króla Opaliński kazał straży otoczyć siedzibę prymasa i zatoczyć 
przed nią działa. Uchański więźnia wydał, co więcej, musiał przybyć na Wawel 
i pokornie prosić Henryka o przebaczenie.`,
		`Stanisław Grzybowski, "Dzieje Polski i Litwy (1506-1648)", Wielka Historia Polski t.4, Kraków 2000, str. 180`,
	},
	"07": {
		`Polska korzystała wtedy z błogosławionych skutków swego własnego, przez 
sześćset lat chrześcijańskiej już historii wyrobionego stosunku do spraw wiary.
Wybierano dysydentów, bo to byli ludzie najbardziej umysłowo rozbudzeni, czynni,
odważni, mający słuszny program. Zalecały ich sejmikom względy po ziemsku rzeczowe,
a nie wyznaniowe. Dawano Rejowi mandat, ponieważ był zdolnym, rozumnym i 
przyzwoitym człowiekiem. Jego kalwinizm ani go degradował, ani protegował.`,
		`Paweł Jasienica, "Polska Jagiellonów", 1965, str. 375-376`,
	},
	"08": {
		`Szlachta nie wyhodowała ani nowych odmian zbóż, ani nawet polskich ras koni,
krów, świń czy owiec. "Potrafiła głównie produkować wódkę, która przynosiła
jej zawsze znaczny i pewny dochód" - komentował jeden z najwybitniejszych
polskich historyków wsi, Stefan Inglot (1902-1994).`,
		`Adam Leszczyński, "Ludowa historia Polski", 2020, str. 134`,
	},
	"09": {
		`Po zamknięciu sejmu koronacyjnego Batory, na czele swojej węgierskiej piechoty,
ruszył natychmiast do Warszawy. W Rawie spotkał się z Opalińskim, za nim pospieszyli 
inni senatorowie wielkopolscy, nawet Czarnkowski stawił się pokornie. Uchański
wymawiał się chorobą; król uprzejmie odparł, że wobec tego sam do niego ze swoimi
Węgrami zajedzie, więc prymas pospiesznie wyzdrowiał.`,
		`Stanisław Grzybowski, "Dzieje Polski i Litwy (1506-1648)", Wielka Historia Polski t.4, Kraków 2000, str. 192`,
	},
	"10": {
		`Ziemiański ekonomista kalkulował w sposób całkowicie bezwzględny. Właścicielowi
opłacało się mieć jak największą liczbę poddanych żyjących na tyle dostatnio,
żeby nie kradli  i nie zbiegli, ale równocześnie nie na tyle bogato, żeby mieli
z czego akumulować. Jeżeli gromadzili bogactwo, to siłą rzeczy działo się 
to kosztem właściciela, który nie ściągnął z nich  wszystkich nadwyżek: 
zamożność chłopska świadczyła o tym, że szlachcic nie eksploatuje swojego
majątku w pełni.`,
		`Adam Leszczyński, "Ludowa historia Polski", 2020, str. 133`,
	},
	"11": {
		`W 1503 roku posłowie króla polskiego zastali bramy Warszawy zamknięte na głucho,
a na murach tłumy pospólstwa, które w naczystszej mowie polskiej lżyło 
i bezcześciło wszystkich Polaków. Odpowiedź była spokojna, całkiem logiczna, 
brzmiała zaś mniej więcej tak: Polakom wymyślacie? A wy sami co za jedni jesteście?
Jeszcze w tym samym XVI stuleciu Warszawa miała stać się sercem i stolicą 
Rzeczypospolitej Obojga Narodów.`,
		`Paweł Jasienica, "Polska Jagiellonów", 1965, str. 280`,
	},
	"12": {
		`Nie trzeba było nawet politycznej zręczności, tylko cierpliwego czekania,
by doprowadzić do wcielenia - zgodnie z traktatem krakowskim - Księstwa
Pruskiego do Korony, co z pewnością zmieniłoby sytuację geopolityczną w 
naszym regionie Europy na całe wieki. W 1563 roku król Zygmunt August
zgodził się na rozszerzenie dziedziczenia w Księstwie Pruskim o branderburską
linię Hohenzollernów. Tutaj dopiero Stańczyk miałby powody do snucia jak
najgorszych przypuszczeń, tym bardziej że już w 1539 roku Zygmunt Stary,
odrzucając podobne starania Albrechta, wskazywał, że Korona nie może
dopuścić - nawet w najdalszej perspektywie - do połączenia w jednym ręku
Berlina i Królewca.`,
		`Grzegorz Kucharczyk, "Prusy. Pięć wieków", 2020`,
	},
}
