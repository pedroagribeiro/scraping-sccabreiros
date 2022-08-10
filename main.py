import pycurl
from io import BytesIO
from bs4 import BeautifulSoup
import pandas as pd

b_obj = BytesIO()
crl = pycurl.Curl()
crl.setopt(crl.URL,'https://www.zerozero.pt/edition.php?id_edicao=156405')
crl.setopt(crl.WRITEDATA, b_obj)
crl.perform()
crl.close()
get_body = b_obj.getvalue()
body = get_body.decode('windows-1252')

soy = BeautifulSoup(body, 'html.parser')
table = soy.find('div', {'id':'edition_table'})

body = table.find('tbody')

df = pd.DataFrame(columns=['pos','none1','none2','pontos','jogos','vit','emp','der','gm','gs','dg','none3','equipa'])


for x in body:
    arr = []
    team = x.find('td',{'class':'text'})
    for y in x:        
        arr.append(y.string)
    arr.append(team.find('a').string)
    df.loc[len(df)] = arr
    
df_f = df.drop(['none1','none2','none3'],axis=1)

print(df_f)