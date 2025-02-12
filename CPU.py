class CPU:
        def __init__(self):
                self.regitry = []
                self.registrySize = 0
                self.registryMaxSize = 30
                self.registryPosition = 0
                self.fillRegistry()
        
        def fillRegistry(self):
                for i in range(self.registryMaxSize):
                        self.regitry[i] = 0
                        
        def OPCodeReader(self, OPCode):
                pass