grammar JSON;

json: value;
value: object | array | str | num | boolean | null;
object: LEFT_BRACKET RIGHT_BRACKET | LEFT_BRACKET members RIGHT_BRACKET;
members: member | member COMMA member;
member: str COLORN value;
array:  LEFT_SQUARE_BRACKET RIGHT_SQUARE_BRACKET | LEFT_SQUARE_BRACKET elements RIGHT_SQUARE_BRACKET;
elements: value | value COMMA value; 
str: DOUBLE_QUOTE LETTER DOUBLE_QUOTE | DOUBLE_QUOTE DOUBLE_QUOTE;
num: NUMBER | MINUS NUMBER;
boolean: TRUE | FALSE;
null: NULL;

LEFT_BRACKET: '{';
RIGHT_BRACKET: '}';
LEFT_SQUARE_BRACKET: '[';
RIGHT_SQUARE_BRACKET: ']';
COMMA: ',';
COLORN: ':';
TRUE: 'true';
FALSE: 'false';
NULL: 'null';
DOUBLE_QUOTE: '"';
WHITE_SPACE : [ \r\t\n]+ -> skip;
LETTER: [a-zA-Z0-9]+;
MINUS: '-'; 
NUMBER: [0-9]+;
