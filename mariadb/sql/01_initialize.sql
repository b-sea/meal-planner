CREATE DATABASE IF NOT EXISTS `meal_planner`;

USE `meal_planner`;

CREATE TABLE IF NOT EXISTS `plans` (
    `id` UUID PRIMARY KEY,
    `min_kcal_target` FLOAT NOT NULL,
    `max_kcal_target` FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS `days` (
    `id` UUID PRIMARY KEY,
    `date` DATE NOT NULL,
    `plan_id` UUID NOT NULL,
    
    CONSTRAINT fk_day_plan_id FOREIGN KEY(`plan_id`) REFERENCES `plans`(id),
);

CREATE TABLE IF NOT EXISTS `meals` (
    `id` UUID PRIMARY KEY,
    `name` NVARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `days_meals` (
    `day_id` UUID NOT NULL,
    `meal_id` UUID NOT NULL,

    PRIMARY KEY(`day_id`, `meal_id`),
    
    CONSTRAINT fk_day_meal_day_id FOREIGN KEY(`day_id`) REFERENCES `days`(id),
    CONSTRAINT fk_day_meal_day_id FOREIGN KEY(`meal_id`) REFERENCES `meals`(id)
);

CREATE TABLE IF NOT EXISTS `units` (
    `id` UUID PRIMARY KEY,
    -- TODO: store unit data properly
);

CREATE TABLE IF NOT EXISTS `nutrition` (
    `id` UUID PRIMARY KEY,
    `kcal` FLOAT NOT NULL,
    `total_fat` FLOAT NOT NULL,
    `saturated_fat` FLOAT NOT NULL,
    `trans_fat` FLOAT NOT NULL,
    `sodium` FLOAT NOT NULL,
    `total_carbs` FLOAT NOT NULL,
    `fiber` FLOAT NOT NULL,
    `total_sugars` FLOAT NOT NULL,
    `protein` FLOAT NOT NULL,
);

CREATE TABLE IF NOT EXISTS `food` (
    `id` UUID PRIMARY KEY,
    `name` NVARCHAR(255) NOT NULL,
    `serving_size` FLOAT NOT NULL,
    `unit_id` UUID NOT NULL,
    `dash_group` INT NOT NULL,
    `nutrition_id` UUID NOT NULL,

    CONSTRAINT fk_food_unit_id FOREIGN KEY(`unit_id`) REFERENCES `units`(id),
    CONSTRAINT fk_food_nutrition_id FOREIGN KEY(`nutrition_id`) REFERENCES `nutrition`(id)
);

CREATE TABLE IF NOT EXISTS `ingredients` (
    `id` UUID PRIMARY KEY,
    `quantity` FLOAT NOT NULL,
    `unit_id` UUID NOT NULL,
    `meal_id` UUID NOT NULL,
    `order` INT NOT NULL,

    CONSTRAINT fk_ingredient_unit_id FOREIGN KEY(`unit_id`) REFERENCES `units`(id),
    CONSTRAINT fk_ingredient_meal_id FOREIGN KEY(`meal_id`) REFERENCES `meals`(id)
);