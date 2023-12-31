# Factory method (фабричный метод)
**Фабричный метод** — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе, 
позволяя подклассам изменять тип создаваемых объектов.

Он используется для создания объектов, скрывая детали конкретного 
процесса создания и делегируя это создание подклассам. 
Factory Method позволяет определить интерфейс для 
создания объекта, но оставляет решение о конкретном 
классе объекта для подклассов.

Используя фабрику, вы можете создавать объекты, не заботясь о деталях их создания, и в будущем легко добавлять новые 
реализации объектов, не изменяя существующий код.