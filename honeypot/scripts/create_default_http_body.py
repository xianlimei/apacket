#!/usr/bin/env python
# -*- coding: UTF-8 -*-

fd = open('default.html', 'r')

lines = fd.readlines()

body = "#!/usr/bin/env python\n# -*- coding: UTF-8 -*-\n\nhttp_body='"
for line in lines:
    line = line.strip()
    body = '%s%s\\n ' % (body, line)
fd.close()

body = "%s'" % body

fd = open('./http_default_body.py', 'w')
fd.write(body)
fd.close
