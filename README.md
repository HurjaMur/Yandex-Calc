# Предисловие:

Я не мастак писать душевные тексты, но позволю себе кое-что написать. Этот вайтишный курс был моим первым, и я надеюсь, не последним, потому что... Мне зашло, да, действительно, зашло. 
Да, этот курс не идеальный, далеко не идеальный. Особенно мне, как новичку в айти, было тяжело понимать, хоть авторы и пытались сделать курс для школьников, но, будем честны, у них не очень получилось. 
Было ОЧЕНЬ много моментов, когда половину, а то и больше информации просто не объясняли: мол, гуглите, ищите. Может, это правильно, типа, вайтишники должны всё гуглить, искать, но тогда для чего этот курс?
Школьная система построена на постулате объяснения учителей; чаще всего, в неё не входит момент самообразования и, тем самым, школьников, по большей части, лишают возможности искать информацию, это просто не нужно.
Этот курс полностью противоположен: в нём без самообразования – никуда.
Может, всё просто: курс создан, чтобы объяснить базу школьникам, а дальше они сами? Не знаю. Может, всё, что я здесь написал – полный бред, и авторы задумали всё совсем по-другому? Не знаю. Оставлю эти вопросы открытыми.

P.S.Как быстро год пролетел, да?

# Гайд по запуску проекта

Чтобы запустить проект:                                                                    
$ git clone https://github.com/HurjaMur/Yandex-Calc.git                                                 
$ go add github.com/mattn/go-sqlite3                                                             
$ docker build -t calc:v1 . (ждём пока забилдится)                                                      
$ docker run -p 8000:8000 calc:v1    

#Что в проекте реализовано
