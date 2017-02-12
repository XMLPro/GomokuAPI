from enum import Enum


class Status(Enum):
    """盤面の目の状態を表します。

    white - 白石が打たれている。
    black - 黒石が打たれている。
    none  - 何も打たれていない。
    """
    white = 1
    black = 2
    none = 3


class Gomoku:
    """五目並べクラスです。

    盤面は配列で持っています。
    盤面への打ち込み、ゲーム終了判定等はこのクラスから行われます。
    """
    def __init__(self, x_size=15, y_size=15):
        """盤面の初期化を行います。

        x_size - 盤面の横幅
        y_size - 盤面の縦幅
        """
        self.board = [[Status.none for x in range(x_size)] for y in range(y_size)]


    def _check_0_degrees():
        """横方向(ー)で5目揃っているかを確認します。
        
        return - 揃っている -> true
        """
        return False


    def _check_45_degrees():
        """斜め方向(／)で5目揃っているかを確認します。
        
        return - 揃っている -> true
        """
        return False


    def _check_90_degrees():
        """縦方向(｜)で5目揃っているかを確認します。
        
        return - 揃っている -> true
        """
        return False


    def _check_135_degrees():
        """斜め方向(＼)で5目揃っているかを確認します。
        
        return - 揃っている -> true
        """
        return False


    def put(status, x, y):
        """碁石を置きます。

        status - 石の色: Status.xxxx
        x      - x座標
        y      - y座標
        return - 置けた -> true
        """
        return False
