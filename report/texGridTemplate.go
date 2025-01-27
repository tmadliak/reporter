package report

const defaultGridTemplate = `
%use square brackets as golang text templating delimiters
\documentclass[a3paper]{article}
\usepackage{graphicx}
\usepackage[margin=0.1in]{geometry}
\usepackage{xcolor} 
\usepackage{helvet}


\definecolor{background}{RGB}{25,27,30} 
\definecolor{text}{RGB}{255,255,255} 

\pagecolor{background}
\color{text}
 \renewcommand{\familydefault}{\sfdefault}

\graphicspath{ {images/} }
\begin{document}
\title{[[.Title]] [[if .VariableValues]] \\ \small Report variables:[[.VariableValues]] [[end]]}
\date{Report period: [[.FromFormatted]] to [[.ToFormatted]]}
\maketitle
\begin{center}
[[range .Panels]][[if .IsPartialWidth]]\begin{minipage}{[[.Width]]\textwidth}
\includegraphics[width=\textwidth]{image[[.Id]]}
\end{minipage}
[[else]]\par
\vspace{0.2cm}
\includegraphics[width=\textwidth]{image[[.Id]]}
\par
\vspace{0.2cm}
[[end]][[end]]

\end{center}
\end{document}
`
