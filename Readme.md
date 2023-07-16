# check-certs

### Описание
Утилита для получения сертификатов из vault, проверки их, а так же отправки запорса на перевыпуск сертификата.

### Конфигурирование
Утилита по умолчанию ищет свой конфиг:
1. Рядом с исполняемым файлом
2. В каталоге ./.k8s/ относительно исполняемого файла

При необходимости можно запустить с параметром: `config-custom-folder` и передать нужный каталог

Можно воспользоваться установкой переменных окружения:
- CHECK_CERTS_VAULT_ADDRESS=http://127.0.0.1:8200
- CHECK_CERTS_VAULT_CLIENT_TIMEOUT=60s
- CHECK_CERTS_VAULT_TOKEN=TOKEN
- CHECK_CERTS_VAULT_MOUNT_PATH=secret
- CHECK_CERTS_VAULT_CERT_PATHS=path1 path2

При конфигурировании через переменные окружения конфигурационный файл все равно ищется и испоользуется для установки значений по умолчанию

### Дополнительные возможности
Запустить с ключем `version` для получения сведений о версии билда
Запустить с ключем `local-run` для использования файла конфигурации `config_local.yaml` вместо `cconfig.yaml`