# การ setup project

1.clone project
``` shell
git clone https://github.com/fozaza/OmaChan-api.git \
cd OmaChan-api
```

2.Download Docker
- arch base
``` sudo pacman -S docker ```
ถ้ามี yay packmanager
``` yay -S docker ```

- debian base or ubuntu base
``` sudo apt install docker ```

- windows [download here](https://www.docker.com/)

3.Download Docker compose
- archlinux base
``` sudo pacman -S docker-compose ```
ถ้ามี yay packmanager
``` yay -S docker-compose ```

- debian base or ubuntu base
``` sudo apt install docker-compose-plugin


4.build image
 ```docker build -t omachan-image -f ./Dockerfile . ```
เพิ่มเติมสามารถ config file docker เพื่อเปลี่ยน password email(ใส่เป็น email อะไรก็ได้ยังไม่สามารถใช้งาน email)
หา
```
 ENV EMAIL = "test@email.com"
 ENV PASSWORD = "12345678"
```

เพิ่มเติมสามารถ config file docker-compose เพื่อเปลี่ยน password databaseName user
ตรง environment  progres 
```
 POSTGRES_DB: oma_chan_data # <<< Database name
 POSTGRES_USER: root # <<< User name
 POSTGRES_PASSWORD: qqee22rr43 # <<< Password
```

6.run container
``` docker compose up -d ```

วิิธีปิดการใช้งาน
``` docker compose down ```

ในกรณีที่มีการ config file new เเละเคย run docker ตัองลบ volum เก่าทิ่งก่อน

net step
- [web]()
- [esp32]()
