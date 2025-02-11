import json

file_path = "./data/tests.json"

with open(file_path, "r", encoding="utf-8") as file:
    data = json.load(file)


for item in data:
    item["input_for_user"] = item["input"]
    item["output_for_user"] = item["output"]

# Преобразуем обновленный массив обратно в JSON
updated_json = json.dumps(data, indent=2)
print(updated_json)
with open("./test/test.json", "w", encoding="utf-8") as file:
    # Записываем данные в файл
    json.dump(updated_json, file, indent=2, ensure_ascii=False)