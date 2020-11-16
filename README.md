# СЕРВИС ДЛЯ БЫСТРОГО ПОИСКА АНАГРАММ В СЛОВАРЕ

## Оглавление:

**[ЧАСТЬ 0: Описание проекта](#description)**  
**[ЧАСТЬ 1: Формат данных](#format)**    
**[ЧАСТЬ 2: API](#api)**  
&emsp;**[2.1 Получение всех анаграмм из словаря по заданному слову](#get)**  
&emsp;**[2.2 Загрузка словаря](#load)**         


**** 
<a name="description">  </a>

### ЧАСТЬ 0. Описание проекта

Сервис позволяет быстро найти анаграммы в словаре для заданного слова.
Слова могут включать только буквы русского и английского алфавита.

Сервис предоставляет следующие API:
1) API для получения всех анаграмм по заданному слову.
2) API для загрузки словаря.

****
<a name="format">  </a>

### ЧАСТЬ 1: Формат данных

**ФОРМАТЫ ВХОДНЫХ И ВЫХОДНЫХ ДАННЫХ:**   

API принимает и возвращает данные в формате ```JSON```

``` 

{
	data                        - данные;           
	error (bool)                - есть ли ошибка;          
	errorText (string)          - название ошибки;
	additionalErrors (struct)   - дополниние к ошибкам;
}
 
```
****
<a name="api">  </a>

### ЧАСТЬ 2: API  

<a name="get">  </a>

### **2.1 Получение всех анаграмм из словаря по заданному слову** 

**URL:**  
```http://localhost:8080/get?word=foobar```

**METHOD:**  
```GET```

**PARAMS:**   
```
Required:  
 word - string, слово для поиска; 
```

**RESPONSE BODY:**  

```
{
    "data": ["foobar","boofar"],
    "error": false,
    "errorText": "",
    "additionalErrors": null
}
```
<a name="load">   </a>

### **2.2 Загрузка словаря**  

**URL:**  
```http://localhost:8080/load```

**METHOD:**  
```POST```

**PARAMS:**   
```
None 
```

**REQUEST BODY:**
```
["foobar", "aabb", "baba", "boofar", "test"]
```

**RESPONSE BODY:**  
```
{
    "data": "Words uploaded successfully",
    "error": false,
    "errorText": "",
    "additionalErrors": null
}
```
