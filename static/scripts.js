function loadTodayContent() {
  // Отправляем GET-запрос к серверу для получения содержимого "Сегодня"
  fetch('/today')
      .then(response => response.text())
      .then(data => {
          // Заменяем содержимое элемента main-content на загруженные данные
          document.getElementById('main-content').innerHTML = data;
      })
      .catch(error => console.error('Ошибка при загрузке раздела "Сегодня":', error));
}
