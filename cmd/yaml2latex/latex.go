package main

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

\title{XVI wiek\\ \noindent{Kalendarium}}
\author[Piotr Jaskulski]{Piotr Jaskulski}
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

\usepackage{makeidx}
\makeindex

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

\par Opisy wydarzeń są dostępne na licencji Creative Commons Uznanie 
Autorstwa 4.0 Międzynarodowe.
\par\textit{Plik wygenerowany w \monthyear}
\end{fullwidth}

\cleardoublepage

\mainmatter
`

var latex_end = `
\backmatter

\printindex

\end{document}
`
