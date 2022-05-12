-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Май 12 2022 г., 19:35
-- Версия сервера: 8.0.24
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `events_backend`
--

-- --------------------------------------------------------

--
-- Структура таблицы `events`
--

CREATE TABLE `events` (
  `id` int NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `shortDescription` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `eventDate` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `latitude` double NOT NULL,
  `longitude` double NOT NULL,
  `preview` varchar(255) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `events`
--

INSERT INTO `events` (`id`, `title`, `shortDescription`, `description`, `eventDate`, `latitude`, `longitude`, `preview`) VALUES
(1, 'Party', 'Alcohol', '10 liters of vodka', 'May 15 09:00', 49.834528, 23.981823, 'https://d1csarkz8obe9u.cloudfront.net/posterpreviews/party-night-banner-design-template-9c9d0ee04d93c1e567f473f6a683c61c_screen.jpg?ts=1610236562'),
(2, 'Party', 'Alcohol', '20 liters of vodka', 'May 18 09:00', 48.834528, 22.981823, 'https://d1csarkz8obe9u.cloudfront.net/posterpreviews/party-night-banner-design-template-9c9d0ee04d93c1e567f473f6a683c61c_screen.jpg?ts=1610236562');

-- --------------------------------------------------------

--
-- Структура таблицы `images`
--

CREATE TABLE `images` (
  `id` int NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `event_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `images`
--

INSERT INTO `images` (`id`, `url`, `event_id`) VALUES
(1, 'https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F184375039%2F474927372937%2F1%2Foriginal.20211111-155142?w=600&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C236%2C4724%2C2362&s=4e575120dfe7909c04845ca84c81491d', 1),
(2, 'https://whatdreammeans.com/wp-content/uploads/2022/02/party.jpg', 1),
(3, 'https://www.adiyprojects.com/wp-content/uploads/2021/12/Party-1.jpg', 1),
(4, 'https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F184375039%2F474927372937%2F1%2Foriginal.20211111-155142?w=600&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C236%2C4724%2C2362&s=4e575120dfe7909c04845ca84c81491d', 2),
(5, 'https://whatdreammeans.com/wp-content/uploads/2022/02/party.jpg', 2),
(6, 'https://www.adiyprojects.com/wp-content/uploads/2021/12/Party-1.jpg', 2);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `events`
--
ALTER TABLE `events`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `images`
--
ALTER TABLE `images`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `events`
--
ALTER TABLE `events`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT для таблицы `images`
--
ALTER TABLE `images`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
