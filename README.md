# CLIFileDirectory

En liten CLI-applikasjon for å jobbe litt med `os` og fil-navigasjon i Go.

Prosjektet er fortsatt ganske *bare-bones* og uferdig, men det fungerer å navigere i filsystemet direkte fra terminalen, og man blir `cd`-et inn i mappen man står i når programmet avsluttes.

---

## Funksjoner

- Piltast-navigasjon
- Gå inn og ut av mapper
- Starter i mappen der programmet kjøres

---

## Nødvendigheter

- Go **1.20+**

---


## For å bruke

1. Klon repoet
2. I root-mappen til repoet:
```shell

go install ./cmd/clifiledirectory
```
```shell

mv ~/go/bin/clifiledirectory ~/go/bin/filenav
```
3. Sjekk:
```shell

echo $PATH | grep go/bin
```

4. Hvis ingenting printer, legg til dette i din shell-config (ellers hopp over dette):
```shell

export PATH="$HOME/go/bin:$PATH"
```

5. 
```shell

nano ~/.zshrc #eller ~/.bashrc
```
6. legg til denne shell-funksjonen nederst:
```shell

navcd() {
     local dir
     dir="$(command filenav)"
     if [[ -n "$dir" && -d "$dir" ]]; then
     cd "$dir"
     fi
}
```

7. Reload shell:
```shell

source ~/.zshrc #eller source ~/.bashrc
```
8. Kjør:
```shell

navcd
```
