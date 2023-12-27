# URL checker

## Steps

1. Dato un file `txt` contenente un URL su ogni riga, andare a chiamare quell'URL e tracciare lo status code di risposta.
1. Iniziare con un'implementazione sincrona.
1. Cambiare l'implementazione da sincrona ad asincrona. (lasciare l'opzione di run in sincrono tramite i flags).
1. tracciare anche l'elapsed time.
1. prevedere un flag che mi decida se creare un file per ogni url o un file globale per tutti gli url
1. creare un file diverso per ogni giorno che il programma gira (oggi creare file con prefisso `2023_12_18`, domani `2023_12_19`)
1. tutto quello che non Ã¨ specificato, decidetelo voi!

// TODO: expand this file and provide the user with detailed information on how to run the program.

### Usage
1. Write the urls to check in the file urls.txt under the cmd/urlschecker folder
1. Execute the program to run it in the default mode which creates a Tmp folder in your current path containing a txt file with the status code of every url contained in the urls.txt file.
1. Default Mode runs in async mode, checking all the urls at the same time if u want to run the program in sync mode (checking one url at a time pass the -s flag as an argument)
1. If u want to print the output to your console instead of on a file use the flag -c as an argument
1. If u want to write a single file for every url checked use the flag -m as an argument, the program will create subfolder under the Tmp folder containing every single file created for each url.
