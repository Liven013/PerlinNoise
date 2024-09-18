Развитие современных компьютерных систем дает возможность использовать все больше и больше ресурсов для создания и демонстрации визуальной информации, что дает возможность увеличивать реалистичность объектов и повышает разнообразие вариантов итогового изображения. Одним из инструментов при создании подобного контента является процедурная генерация, которая позволяет не хранить огромные объемы данных, а воспро-изводить их с помощью алгоритмов генерации по мере необходимости. Большую популяр-ность в игровой индустрии получил алгоритм шума, разработанный Кимом Перлином.
Этот алгоритм применяют при генерации различных ландшафтов, рельефов и текстур, что говорит об универсальности метода. В данной статье будет представлено исследование зависимости визуальных характеристик полученных изображений от параметров, использу-емых при генерации шума.
Для создания шума на исходное изображение накладывается сетка, для каждого узла которой определяется случайный градиент. Далее значение шума рассчитывается для внут-ренних точек. Для этого определяются векторы от 4 ближайших вершин сетки до точки, в которой рассчитывается шум, и вычисляются скалярные произведения векторов и соответ-ствующих градиентов. Итоговый результат значения шума получается интерполяцией полу-ченных скалярных произведений.
