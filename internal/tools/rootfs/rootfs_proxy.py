import ctypes
from typing import List

POLICY_LIBRARY = "./librootfs.so"


class SecurityPolicyProxy:
    def __init__(self, policy_lib: str):
        self.lib = ctypes.cdll.LoadLibrary(policy_lib)
        self.lib.get_rootfs_layer_hashes.argtypes = [
            ctypes.c_char_p,
            ctypes.c_char_p,
            ctypes.c_char_p,
            ctypes.POINTER(ctypes.c_int),
            ctypes.POINTER(ctypes.POINTER(ctypes.c_char_p)),
        ]
        self.lib.get_rootfs_layer_hashes.restype = None

        self.lib.clean_up.argtypes = [
            ctypes.c_int,
            ctypes.POINTER(ctypes.c_char_p),
        ]
        self.lib.clean_up.restype = None

    def generate_policy_agent(self, image: str, tag: str) -> List[str]:
        c_size = ctypes.c_int(0)
        c_layers = ctypes.POINTER(ctypes.c_char_p)()
        self.lib.get_rootfs_layer_hashes(
            image.encode("utf-8"),
            tag.encode("utf-8"),
            "local".encode("utf-8"),
            ctypes.byref(c_size),
            ctypes.byref(c_layers),
        )
        layers = []
        for i in range(c_size.value):
            layers.append(c_layers[i].decode("utf-8"))

        self.lib.clean_up(c_size, c_layers)

        return layers


spp = SecurityPolicyProxy(POLICY_LIBRARY)
res = spp.generate_policy_agent("rust", "1.52.1")
for layer in res:
    print(layer)
