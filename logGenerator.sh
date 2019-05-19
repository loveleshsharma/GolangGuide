#!/bin/bash

arr=('68.99.197.11 - Reichert4658 772 [2019-05-14T15:55:15+05:30] PATCH /robust/communities/magnetic/leverage 503 21888' '110.117.105.146 - Hilll3246 961 [2019-05-14T15:55:15+05:30] DELETE /bricks-and-clicks/world-class/roi/e-services 302 10226' '95.210.66.29 - Block7568 553 [2019-05-14T15:55:15+05:30] PATCH /enhance/b2b 400 27215' 
'120.46.218.22 - Nader8357 124 [2019-05-14T15:55:15+05:30] POST /drive/transform/integrated/disintermediate 205 1449' 
'85.190.68.55 - Aufderhar8452 652 [2019-05-14T15:55:15+05:30] PATCH /models/wireless 503 22614' 
'34.231.156.74 - Lind8153 184 [2019-05-14T15:55:15+05:30] DELETE /engage/mindshare 201 15215' 
'52.36.129.86 - Heller1083 869 [2019-05-14T15:55:15+05:30] PUT /convergence/markets/grow 100 26569' 
'127.156.140.115 - Wilderman8420 338 [2019-05-14T15:55:15+05:30] POST /evolve/iterate/users 401 9524' 
'69.66.18.232 - Bartell4503 940 [2019-05-14T15:55:15+05:30] HEAD /visualize/relationships 406 14682' 
'160.33.29.28 - Hauck6761 791 [2019-05-14T15:55:15+05:30] PUT /cross-media/functionalities/infrastructures/mindshare 401 3020' 
)

rm log.txt
for i in {1..5}
do
    for j in {1..1000}
    do
        echo  ${arr[$(( $RANDOM % 10 + 1 ))]}>> log.txt
    done
    sleep 1
done