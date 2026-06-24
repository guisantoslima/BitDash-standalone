
"""
BitDash Analytics Engine.

Responsável por cálculos de dashboard, agregações mensais, posição por
ativo e exportações. Acionado pelo host Go via subprocess + JSON, conforme
Blueprint seção 14 (Estratégia de integração Go <-> Python).

Implementação completa dos cálculos ocorre no Sprint 3.
"""

__version__ = "0.1.0"

