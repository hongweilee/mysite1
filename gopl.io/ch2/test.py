#coding=utf-8
class Home():
    driver=222
    @classmethod
    def setUpClass(cls):
        cls.driver=1111

class Element():
    @classmethod
    def systmeRole(cls):
        Home().setUpClass()
        driver = Home.driver
        print driver



Element().systmeRole()
