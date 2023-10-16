# Singleton
**Singleton (Одиночка)** - это порождающий паттерн проектирования, который гарантирует, что класс имеет только один 
экземпляр, и предоставляет глобальную точку доступа к этому экземпляру. Это полезно, когда есть необходимость в 
одном общем объекте, который управляет общими ресурсами, такими как конфигурация, 
пул соединений к базе данных, журналы и т.д.

Использование Singleton позволяет создать общий объект, который может быть разделяемым между разными частями программы, 
и обеспечивает доступ к нему из любой точки приложения.