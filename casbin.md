### Casbin example
try paste this code above to this [casbin editor](https://casbin.org/editor/)

---

**Model**
```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && (r.obj == p.obj || p.obj == "*") && (r.act == p.act || p.act == "*")
```

---

**Policy**
```
p, superadmin, *, *
p, admin, product, *
p, admin, transaction, *
p, cashier, product, read
p, cashier, transaction, create
p, cashier, transaction, read
p, cashier, transaction, send-summary

g, root, superadmin
```

---

**Request**
```
superadmin, product, read
superadmin, transaction, delete
superadmin, whatever, whatever
root, product, read
root, transaction, delete
root, whatever, whatever
admin, product, whatever
admin, product, delete
admin, transaction, send-summary
admin, whatever, read
cashier, product, create
cashier, product, read
cashier, transaction, send-summary
anonymous, product, read
```

---

**Enforcement Result (expected)**
```
true
true
true
true
true
true
true
true
true
false
false
true
true
false
```

