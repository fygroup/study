//--json-------------------------
import net.sf.json.JSONObject;
1、JSON字符串->json对象->java对象
JSONObject jsonobject = JSONObject.fromObject(jsonStr);
MyObject mc = (MyObject)JSONObject.toBean(jsonobject,MyObject.class);
User user=(User)JSONObject.toBean(jsonobject,User.class);
2、map对象 -> json对象 -> json string
Map<String,String> map = new Map<String,String>(); //new Map<>()也行
map.put("a","1");
JSONObject jsonobject = JSONObject.fromobject(map);
String jsonstr = jsonobject.toString();








