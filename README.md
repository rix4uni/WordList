# WordList

custom fuzzing wordlist `fuzzing_list.txt`
```
cat urls.txt | sed 's|\(.*\)/[^/]*$|\1|' | cut -d"/" -f4,5,6,7,8,9,10,11 | tr "/" "\n" | sed '/^$/d' | anew fuzzing_list.txt
```

custom dns wordlist `dns-wordlist.txt`
```
cat alltargets.txt | sed 's/\.[^.]*$//' | tr "." "\n" | egrep -v '^[0-9]*$' | anew dns-wordlist.txt
```

custom urls for nuclei automation
```
cat urls.txt | sed 's|\(.*\)/[^/]*$|\1|' | cut -d'/' -f1-6 | anew urls.txt
cat urls.txt | sed 's|\(.*\)/[^/]*$|\1|' | cut -d'/' -f1-5 | anew urls.txt
cat urls.txt | sed 's|\(.*\)/[^/]*$|\1|' | cut -d'/' -f1-4 | anew urls.txt
```

`default-username-password.txt`
```
curl -s "https://raw.githubusercontent.com/rix4uni/WordList/main/default-username-password.txt"|cut -d":" -f1 | tee -a username.txt
curl -s "https://raw.githubusercontent.com/rix4uni/WordList/main/default-username-password.txt"|cut -d":" -f2 | tee -a password.txt
```

custom parameters wordlist `params.txt`
```
cat urls.txt | grep "\.php?" | uro | grep "?" | cut -f2 -d"?" | cut -f1 -d"=" | sed '/^\s*$/d'| anew params.txt
```

custom fuzzing wordlist `dicc-onelistforallmicro.txt`
```
curl -s https://raw.githubusercontent.com/maurosoria/dirsearch/master/db/dicc.txt | anew dicc-onelistforallmicro.txt
curl -s https://raw.githubusercontent.com/six2dez/OneListForAll/main/onelistforallmicro.txt | anew dicc-onelistforallmicro.txt
```
