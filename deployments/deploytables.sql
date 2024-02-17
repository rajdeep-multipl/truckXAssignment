CREATE TABLE IF NOT EXISTS `truckxdb`.`sensors` (
  `id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `id_index` (`id` ASC) VISIBLE)

CREATE TABLE IF NOT EXISTS `truckxdb`.`temperatures` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `sensor_id` BIGINT NOT NULL,
  `current_temperature` INT NOT NULL,
  `timestamp` BIGINT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_sensor_id` (`sensor_id` ASC) VISIBLE,
  CONSTRAINT `fk_sensor_id` FOREIGN KEY (`sensor_id`) REFERENCES `truckxdb`.`sensors` (`id`))


CREATE TABLE IF NOT EXISTS `truckxdb`.`aggregated_temperatures` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `max_temperature` INT NOT NULL,
  `min_temperature` INT NOT NULL,
  `avg_temperature` INT NOT NULL,
  `sensor_id` BIGINT NOT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `sensor_id` (`sensor_id` ASC) VISIBLE,
  CONSTRAINT `aggregated_temperatures_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `truckxdb`.`sensors` (`id`))