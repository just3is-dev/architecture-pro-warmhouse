# Key System Flows (TO-BE)

Документ описывает ключевые пользовательские и системные сценарии end-to-end.  
Каждый сценарий фиксирует участников и последовательность шагов, включая публикацию событий в шину.

---

## Flow 1. Покупка модуля и активация в доме

### Участники
- User
- API Gateway
- Billing Service
- Payment Provider
- House Service
- Device Management Service
- Message Broker

### Последовательность
1. User оформляет покупку модуля в интерфейсе (через API Gateway).
2. API Gateway вызывает Billing Service (создать подписку/платёж).
3. Billing Service проводит оплату через Payment Provider.
4. Billing Service публикует событие `SubscriptionActivated` в Message Broker.
5. House Service получает `SubscriptionActivated` и создаёт установленный модуль в доме.
6. House Service публикует событие `InstalledModuleCreated`.
7. Device Management Service получает `InstalledModuleCreated` и начинает процесс регистрации устройств (в рамках модуля).

---

## Flow 2. Поступление телеметрии и выполнение сценария

### Участники
- Partner Device
- Telemetry Service
- Automation Service
- Device Management Service
- Message Broker

### Последовательность
1. Устройство отправляет телеметрию в Telemetry Service.
2. Telemetry Service сохраняет запись и публикует `TelemetryReceived`.
3. Automation Service получает `TelemetryReceived`, проверяет правила/сценарии дома.
4. Если условия выполнены, Automation Service публикует `CommandRequested`.
5. Device Management Service получает `CommandRequested` и отправляет команду устройству.
6. Device Management Service публикует `DeviceCommandSent` или `DeviceCommandFailed`.

---

## Flow 3. Самостоятельное подключение устройства пользователем

### Участники
- User
- API Gateway
- House Service
- Device Management Service
- Message Broker

### Последовательность
1. User инициирует подключение устройства (например, вводит serial_number).
2. API Gateway вызывает Device Management Service (зарегистрировать устройство).
3. Device Management Service создаёт устройство и публикует `DeviceRegistered`.
4. House Service (или Device Service — зависит от владения) выполняет привязку устройства к дому и публикует `DeviceAssignedToHouse`.
5. Устройство становится доступным для телеметрии и сценариев.

---

## Flow 4. Отключение модуля из-за прекращения подписки

### Участники
- Billing Service
- House Service
- Device Management Service
- Automation Service
- Message Broker

### Последовательность
1. Billing Service фиксирует отмену/истечение подписки и публикует `SubscriptionCancelled`.
2. House Service получает событие и деактивирует установленный модуль.
3. House Service публикует `InstalledModuleDeactivated`.
4. Device Management Service ограничивает управление устройствами модуля (например, запрет команд).
5. Automation Service отключает сценарии, связанные с данным модулем (опционально).
