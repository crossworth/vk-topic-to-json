### VK Topic to JSON
Save a VK Topic to a JSON backup

####How to use
This is a command line app, you must call it with following arguments

*Unix*
```bash
vk-topic-to-json -email=YOUR_VK_EMAIL -password=YOUR_VK_PASSWORD -group=VK_GROUP_ID -topic=VK_TOPIC_ID
```

*Windows*
```batch
vk-topic-to-json.exe -email=YOUR_VK_EMAIL -password=YOUR_VK_PASSWORD -group=VK_GROUP_ID -topic=VK_TOPIC_ID
```

You can get the group ID and topic ID from a topic URL

```
https://vk.com/topic-203040506_40506070
                     |GROUP  ||TOPIC  |
```

In this case the *group* is **203040506** and the *topic* is **40506070**.
