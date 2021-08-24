### rapidjson
```c++

rapidjson::StringBuffer buffer;
rapidjson::Writer<rapidjson::StringBuffer> writer(buffer);
jsonDoc.Accept(writer);
std::string formatJsonStr = buffer.GetString();
```