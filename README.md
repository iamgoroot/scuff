

populates templates from json file
`go get github.com/iamgoroot/scuff` 
`scuff mytemplate/scuff.json`

```
{
//optional scuff config:

   "scuff": {
     "delim": {
       "left": "[[",        //default: "{{"
       "right": "]]"        //default: "}}"
     },
     "out": "./out",
     "in": "."
   },
//any property you want
   "project": {
       "shortName": "project1"
   }
 }
 ```