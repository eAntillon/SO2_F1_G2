# Manual Tecnico Practica 1

#####  Universidad de San Carlos de Guatemala
#####  Escuela de Ciencias y Sistemas
#####  Sistemas Operativos 2
-------------------------
## Integrantes
## Grupo2
| Nombre                          | Carnet    |
|----------------------------     |-----------|
| Edi Yovani Tomas Reynoso        | 201503783 |
| Evelyn Alejandra Navarro Ozorio | 201902046 |
| Erick Daniel Antillón Chinchilla| 201906552 |
| Edgar Humberto Borrayo Bartolón | 201503702 |
| Josue Alfredo Gonzalez Caal     | 201602489 |


### SYSTEMA DE METAOS 

#### Librerias
Las librerias de golang se encuentran en la siguiente pagina
```
https://pkg.go.dev/std
```

#### _Menu principal_
![image](/assets/Tecnico/Menuprincipal.png)

Este metodo  permite visualizar el metodo principal del programa en consola.

#### _IOTOP_
![image](/assets/Tecnico/run.png)

Este  permite correr la funcion iotop.

![image](/assets/Tecnico/execute.png)

Este  permite ejecutar la funcion iotop.

para visualizar el IOTOP ejecute el siguiente comando
```
sudo apt install iotop 
```

#### _TOP_

![image](/assets/Tecnico/RunTop.png)

Este  permite correr la funcion top.

![image](/assets/Tecnico/ExecuteTop.png)

Este  permite ejecutar la funcion top.


#### _SYSCALL_

![image](https://user-images.githubusercontent.com/71113297/188338735-c52d63d5-b537-4839-8ef1-9d98db2edfa1.png)

Este  permite correr la funcion Syscall.

![image](https://user-images.githubusercontent.com/71113297/188338755-9d39c942-e4ba-4a96-b70e-c1aca3c437b2.png)

Y esto permite devolver los resultados.

![image](https://user-images.githubusercontent.com/71113297/188338779-7c80416c-9a2d-4c04-93a8-b0dab65ccce3.png)

Las funciones que consume para terminar el proceso son las siguientes de syscall

![image](https://user-images.githubusercontent.com/71113297/188338808-59236448-ea99-4286-81e0-36c02dc6a118.png)

#### _MEMSIM_

![image](/assets/Tecnico/executeMemsim.PNG)

Este permite correr la funcion de Memsim

#### _BTIACORA_

Para esta fase se agrega la bitácora que devolverá los usuarios que han utilizado la aplicación.

Para esto se agregan las siguientes funciones:

Leer la bitácora:

![image](https://user-images.githubusercontent.com/71113297/188338942-ab1c3ea7-c7b1-4f20-b352-39cd8062ab62.png)

Al momento que el usuario introduzca su nombre o alias se creará el registro:

![image](/assets/Tecnico/bitacoraUsuarioMemsim.PNG)

Se registrarán los comandos ejecutados:

![image](https://user-images.githubusercontent.com/71113297/188338991-917eef37-9a4b-4104-96f3-38181dfab50b.png)

Almacenará el usuario activo:

![image](https://user-images.githubusercontent.com/71113297/188339009-d03c5a52-4f84-4cd5-a7c5-157bc10fc46a.png)

Y escribirá la bitácora:

![image](https://user-images.githubusercontent.com/71113297/188339036-f90ac87b-551a-4472-a711-22a8520a8e11.png)





## FASE 4

### _BACKEND IMPORTS, METODO POSTLOGIN_
1) Aqui se importan todas las librerias que utilizamos.
2) Utilizamos el metodo PostLogin para verificar que el usuario y contraseña sea el valido
![image](/assets/Tecnico/UltimaFase/backend_1.png)


### _BACKEND METODO POSTREGISTRO, INICIO, MEMSIM_
1) Metodo utilizado para crear un nuevo registro en la base de datos
2) Este metodo lo utilizamos para comprobar que todo este funcionando correctamente, la conexion.
3) Este metodo nos sirve para ejecutar el procedimiento MEMSIM
![image](/assets/Tecnico/UltimaFase/backend_2.png)

### _BACKEND METODO BACKEND_
1) En este metodo nos sirve para tener almacenar todas las rutas de la API
![image](/assets/Tecnico/UltimaFase/backend_3.png)



### _DOCKERFILE DEL BACKEND_
![image](/assets/Tecnico/UltimaFase/backend4_dockerfile.png)

### _DOCKERCOMPOSE DEL BACKEND_
![image](/assets/Tecnico/UltimaFase/backend5_dockercompose.png)

### _DOCKERCOMPOSE DEL FRONTEND_
![image](/assets/Tecnico/UltimaFase/frontend_dockercompose.png)

### _DOCKERFILE DEL FRONTEND_
![image](/assets/Tecnico/UltimaFase/frontend_dockerfile.png)

### _PAGINA HOME_
![image](/assets/Tecnico/UltimaFase/frontend_home.png)

### _PAGINA LOGIN_
![image](/assets/Tecnico/UltimaFase/frontend_login.png)

### _PAGINA DEL REGISTRO_
![image](/assets/Tecnico/UltimaFase/frontend_registro.png)



