# Domain Events

Документ описывает ключевые доменные события целевой системы (TO-BE).
События сгруппированы по сервису-издателю (Publisher), который является владельцем соответствующих данных.

---

## 1. Общие принципы

- Событие публикует сервис-владелец данных (owner).
- Событие — это факт (что уже произошло), а не команда.
- Межсервисное взаимодействие осуществляется через Message Broker.
- Доставка может быть повторной → обработчики должны быть идемпотентными.
- Межсервисная консистентность допускает eventual consistency.

---

## 2. Events by Publisher

---

## 2.1 Telemetry Service

### TelemetryReceived
- Consumers: `Automation Service`
- Назначение: поступило новое измерение от устройства (температура, движение, состояние и т.д.).
- Payload (пример):
    - device_id
    - house_id
    - metric_type
    - value
    - timestamp

---

## 2.2 Device Management Service

### DeviceRegistered
- Consumers: `House Service`
- Назначение: устройство зарегистрировано в системе.

### DeviceAssignedToHouse
- Consumers: `House Service`
- Назначение: устройство привязано к дому.

### DeviceStatusChanged
- Consumers: `Automation Service`, `House Service`
- Назначение: изменилось состояние устройства (online/offline/fault).

### DeviceCommandSent
- Consumers: `Automation Service`
- Назначение: команда успешно отправлена устройству.

### DeviceCommandFailed
- Consumers: `Automation Service`
- Назначение: выполнение команды устройству завершилось ошибкой.

---

## 2.3 House Service

### InstalledModuleCreated
- Consumers: `Device Management Service`
- Назначение: в доме создан установленный модуль.

### InstalledModuleActivated
- Consumers: `Device Management Service`, `Automation Service`
- Назначение: модуль активирован и готов к использованию.

### InstalledModuleDeactivated
- Consumers: `Device Management Service`, `Automation Service`
- Назначение: модуль деактивирован (например, из-за отмены подписки).

---

## 2.4 Automation Service

### CommandRequested
- Consumers: `Device Management Service`
- Назначение: автоматизация инициировала выполнение команды устройству.

### ScenarioCreated
- Consumers: (нет обязательных)
- Назначение: пользователь создал сценарий автоматизации.

### ScenarioEnabled
- Consumers: (нет обязательных)
- Назначение: сценарий активирован.

---

## 2.5 Billing Service

### SubscriptionActivated
- Consumers: `House Service`
- Назначение: подписка на модуль активирована.

### SubscriptionCancelled
- Consumers: `House Service`
- Назначение: подписка отменена или истекла.

---

## 3. MVP Scope

Для MVP обязательными считаются:

- TelemetryReceived
- CommandRequested
- DeviceRegistered
- InstalledModuleCreated
- SubscriptionActivated

Остальные события могут быть добавлены на этапе расширения функциональности.