Unity games aspect ratio patcher (16:9 to 16:10).

An example using Hollow Knight (`~/.steam/steam/steamapps/common/Hollow Knight/hollow_knight_Data/Managed/Assembly-CSharp.dll`):

```shell
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
found expected bytes offset: 13b50
done
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
found expected bytes offset: 13bfe
done
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
found expected bytes offset: 13cfa
done
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
found expected bytes offset: 13d8a
done
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
found expected bytes offset: 2f730b
done
Downloads λ ./unity-ar-patch-16x10 -f Assembly-CSharp.dll
failed to scan file: expected bytes were not found
```
