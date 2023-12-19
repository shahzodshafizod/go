## Tree 1
#### ![en](https://img.shields.io/badge/EN-blue) An address&nbsp;<i>P</i><sub>1</sub> of a record of TNode type is given. The record consists of the <i>Data</i> field (of integer type) and the <i>Left</i> and <i>Right</i> fields (of PNode type). The given record (<i>a tree root</i>) is linked by its <i>Left</i> and <i>Right</i> fields with records of the same type (named the left and right <i>child nodes</i> respectively). Output the Data fields of the tree root and its left and right children. Also output the addresses of left and right child nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан адрес <i>P</i><sub>1</sub> записи типа TNode с полями Data (целого типа), Left и Right (типа PNode &#8212; указателя на TNode). Эта запись (<i>корень дерева</i>) связана полями Left и Right с записями того же типа (левой и правой <i>дочерней вершиной</i>). Вывести значения полей Data корня, его левой и правой дочерних вершин, а также адреса левой и правой дочерних вершин в указанном порядке.

---

## Tree 2
#### ![en](https://img.shields.io/badge/EN-blue) An address&nbsp;<i>P</i><sub>1</sub> of a record of TNode type (a tree root) is given. This record is linked by its <i>Left</i> and <i>Right</i> fields with records of the same type (child nodes); they, in turn, are linked with their own child nodes and so on, until records whose <i>Left</i> and <i>Right</i> fields are equal to NULL. Some of the nodes may have one field (<i>Left</i> or <i>Right</i>) equals NULL. Output the amount of tree nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан адрес <i>P</i><sub>1</sub> записи типа TNode &#8212; корня дерева. Эта запись связана полями Left и Right с другими записями того же типа (дочерними вершинами), они, в свою очередь, &#8212; со своими дочерними вершинами, и так далее до записей, поля Left и Right которых равны NULL (у некоторых вершин может быть равно NULL одно из полей &#8212; Left или Right). Вывести количество вершин дерева.

---

## Tree 3
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty tree and an integer <i>K</i> are given. Output the amount of nodes whose value equals&nbsp;<i>K</i>.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева и число&nbsp;<i>K</i>. Вывести количество вершин дерева, значение которых равно&nbsp;<i>K</i>.

---

## Tree 4
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the sum of values of all tree nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева. Вывести сумму значений всех вершин данного дерева.

---

## Tree 5
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the amount of left child nodes (the tree root should not be counted).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести количество вершин дерева, являющихся левыми дочерними вершинами (корень дерева не учитывать).

---

## Tree 6
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Nodes without children are called <i>the external nodes</i> or <i>the leaf nodes</i> (<i>the leaves</i>). Output the amount of leaf nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. <i>Листом дерева</i> называется его вершина, не имеющая дочерних вершин. Вывести количество листьев для данного дерева.

---

## Tree 7
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the sum of values of all tree leaves.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести сумму значений всех листьев данного дерева.

---

## Tree 8
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a tree is given, the tree contains at least two nodes. Output the amount of tree leaves that are the right child nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева, содержащего по крайней мере две вершины. Вывести количество листьев дерева, являющихся правыми дочерними вершинами.

---

## Tree 9
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. The root node is said to be on the <i>zero level</i>, its child nodes &#8212; on the <i>first level</i>, and so on. Output the <i>depth</i> of the tree (that is, the maximal level of tree nodes). For example, the depth of a tree containing only a root node is equal to&nbsp;0.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Считается, что корень дерева находится на <i>нулевом уровне</i>, его дочерние вершины &#8212; на <i>первом уровне</i> и т.&nbsp;д. Вывести <i>глубину дерева</i>, т.&nbsp;е. значение его максимального уровня (например, глубина дерева, состоящего только из корня, равна&nbsp;0).

---

## Tree 10
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. For each tree level (including the zero one) output the amount of its nodes.  The tree depth is assumed to be not greater than&nbsp;10.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для каждого из уровней данного дерева, начиная с нулевого, вывести количество вершин, находящихся на этом уровне. Считать, что глубина дерева не превосходит&nbsp;10.

---

## Tree 11
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. For each tree level (including the zero one) output the sum of values of its nodes. The tree depth is assumed to be not greater than&nbsp;10.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для каждого из уровней данного дерева, начиная с нулевого, вывести сумму значений вершин, находящихся на этом уровне. Считать, что глубина дерева не превосходит&nbsp;10.

---

## Tree 12
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Using the recursive algorithm named <i>inorder tree walk</i> output the values of all tree nodes as follows: output the left subtree (using inorder tree walk), then output the root node, then output the right subtree (using inorder tree walk too).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести значения всех вершин дерева в <i>инфиксном порядке</i> (вначале выводится содержимое левого поддерева в инфиксном порядке, затем выводится значение корня, затем &#8212; содержимое правого поддерева в инфиксном порядке).

---

## Tree 13
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Using the recursive algorithm named <i>preorder tree walk</i> output the values of all tree nodes as follows: output the root node, then output the left subtree (using preorder tree walk), then output the right subtree (using preorder tree walk too).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести значения всех вершин дерева в <i>префиксном порядке</i> (вначале выводится значение корня, затем содержимое левого поддерева в префиксном порядке, затем &#8212; содержимое правого поддерева в префиксном порядке).

---

## Tree 14
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Using the recursive algorithm named <i>postorder tree walk</i> output the values of all tree nodes as follows: output the left subtree (using postorder tree walk), then output the right subtree (using postorder tree walk too), then output the root node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести значения всех вершин дерева в <i>постфиксном порядке</i> (вначале выводится содержимое левого поддерева в постфиксном порядке, затем &#8212; содержимое правого поддерева в постфиксном порядке, затем &#8212; значение корня).

---

## Tree 15
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree and an integer&nbsp;<i>N</i> (&gt;&nbsp;0) are given. The value of&nbsp;<i>N</i> is not greater than the amount of tree nodes. Output the values of tree nodes whose order numbers are not greater than&nbsp;<i>N</i> (the tree nodes are numbered from&nbsp;1 using inorder tree walk &#8212; see Tree12).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева и число&nbsp;<i>N</i> (&gt;&nbsp;0), не превосходящее количество вершин в исходном дереве. Нумеруя вершины в инфиксном порядке (см. задание Tree12, нумерация ведется от&nbsp;1), вывести значения всех вершин с порядковыми номерами от&nbsp;1 до&nbsp;<i>N</i>.

---

## Tree 16
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree and an integer&nbsp;<i>N</i> (&gt;&nbsp;0) are given. The value of&nbsp;<i>N</i> is not greater than the amount of tree nodes. Output the values of tree nodes whose order numbers are&nbsp;<i>N</i> or greater (the tree nodes are numbered from&nbsp;1 using postorder tree walk &#8212; see Tree14).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева и число&nbsp;<i>N</i> (&gt;&nbsp;0), не превосходящее количество вершин в исходном дереве. Нумеруя вершины в постфиксном порядке (см. задание Tree14, нумерация ведется от&nbsp;1), вывести значения всех вершин с порядковыми номерами от&nbsp;<i>N</i> до максимального номера.

---

## Tree 17
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree and two integers <i>N</i><sub>1</sub>, <i>N</i><sub>2</sub> (0&nbsp;&lt;&nbsp;<i>N</i><sub>1</sub>&nbsp;&lt;&nbsp;<i>N</i><sub>2</sub>) are given. The value of&nbsp;<i>N</i><sub>2</sub> is not greater than the amount of tree nodes. Output the values of tree nodes whose order numbers are in the range&nbsp;<i>N</i><sub>1</sub> to&nbsp;<i>N</i><sub>2</sub> (the tree nodes are numbered from&nbsp;1 using preorder tree walk &#8212; see Tree13).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева и два числа <i>N</i><sub>1</sub>, <i>N</i><sub>2</sub> (0&nbsp;&lt;&nbsp;<i>N</i><sub>1</sub>&nbsp;&lt;&nbsp;<i>N</i><sub>2</sub>), которые не превосходят количество вершин в исходном дереве. Нумеруя вершины в префиксном порядке (см. задание Tree13, нумерация ведется от&nbsp;1), вывести значения всех вершин с порядковыми номерами от&nbsp;<i>N</i><sub>1</sub> до&nbsp;<i>N</i><sub>2</sub>.

---

## Tree 18
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree and an integer&nbsp;<i>L</i> (&#8805;&nbsp;0) are given. Using tree walk of any type (see Tree12&#8722;Tree14) output values of all nodes of the level&nbsp;<i>L</i>. Also output the amount&nbsp;<i>N</i> of these nodes. If the given tree does not contain nodes of level&nbsp;<i>L</i> then output&nbsp;0.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева и неотрицательное число&nbsp;<i>L</i>. Используя любой из описанных в заданиях Tree12&#8722;Tree14 способов обхода дерева, вывести значения всех вершин уровня&nbsp;<i>L</i>, а также их количество&nbsp;<i>N</i> (если дерево не содержит вершин уровня&nbsp;<i>L</i>, то вывести&nbsp;0).

---

## Tree 19
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the maximal value of the tree nodes and the amount of nodes with this value.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести максимальное из значений его вершин и количество вершин, имеющих это максимальное значение.

---

## Tree 20
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the minimal value of the tree nodes and the amount of leaves with this value (the amount may be equal to&nbsp;0).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести минимальное из значений всех его вершин и количество листьев, имеющих это минимальное значение (данное количество может быть равно&nbsp;0).

---

## Tree 21
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the minimal value of its leaves.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести минимальное из значений его вершин, являющихся листьями.

---

## Tree 22
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a tree is given, the tree contains at least two nodes. Output the maximal value of its <i>internal nodes</i> (that is, nodes with children).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева, содержащего по крайней мере две вершины. Вывести максимальное из значений его <i>внутренних вершин</i> (т.&nbsp;е. вершин, не являющихся листьями).

---

## Tree 23
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Using preorder tree walk, find the first tree node with the minimal value and output its address&nbsp;<i>P</i><sub>2</sub>.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести указатель <i>P</i><sub>2</sub> на первую вершину дерева с минимальным значением (вершины перебирать в префиксном порядке).

---

## Tree 24
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Using inorder tree walk, find the last node with the maximal odd value and output its address&nbsp;<i>P</i><sub>2</sub>. If the tree does not contain nodes with odd values then output NULL.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести указатель <i>P</i><sub>2</sub> на последнюю вершину дерева с максимальным нечетным значением (вершины перебирать в инфиксном порядке). Если дерево не содержит вершин с нечетными значениями, то вывести NULL.

---

## Tree 25
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree with <i>N</i>&nbsp;nodes and assign values of the given sequence to tree nodes in order of their creation. Each node of the tree (except for the root) should be a right child. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать дерево из <i>N</i>&nbsp;вершин, в котором каждая вершина (кроме корня) является правой дочерней вершиной. Каждой создаваемой вершине присваивать очередное значение из исходного набора. Вывести указатель на корень созданного дерева.

---

## Tree 26
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree with <i>N</i>&nbsp;nodes and assign values of the given sequence to tree nodes in order of their creation. Each internal node of the tree should have one child: the root has a left child, which has a right child, which has a left child, and so&nbsp;on. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать дерево из <i>N</i>&nbsp;вершин, в котором каждая внутренняя вершина имеет только одну дочернюю вершину, причем правые и левые дочерние вершины чередуются (корень имеет левую дочернюю вершину). Каждой создаваемой вершине присваивать очередное значение из исходного набора. Вывести указатель на корень созданного дерева.

---

## Tree 27
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree with <i>N</i>&nbsp;nodes and assign values of the given sequence to tree nodes in order of their creation. Each internal node of the tree should have one child: an internal node whose value is an odd number has a left child, otherwise it has a right child. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать дерево из <i>N</i>&nbsp;вершин, в котором каждая внутренняя вершина имеет только одну дочернюю вершину, причем если значение вершины является нечетным, то она имеет левую дочернюю вершину, а если четным, то правую. Каждой создаваемой вершине присваивать очередное значение из исходного набора. Вывести указатель на корень созданного дерева.

---

## Tree 28
#### ![en](https://img.shields.io/badge/EN-blue) An even integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree with <i>N</i>&nbsp;nodes; left child nodes of the tree should be leaves, right child nodes should be internal ones. For each internal node create a left child at first, then create a right one (if it exists). Assign values of the given sequence to tree nodes in order of their creation. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано четное число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать дерево из <i>N</i>&nbsp;вершин, в котором каждая левая дочерняя вершина является листом, а правая дочерняя вершина является внутренней. Для каждой внутренней вершины вначале создавать левую дочернюю вершину, а затем правую (если она существует); каждой создаваемой вершине присваивать очередное значение из исходного набора. Вывести указатель на корень созданного дерева.

---

## Tree 29
#### ![en](https://img.shields.io/badge/EN-blue) An even integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree with <i>N</i>&nbsp;nodes. Inner node whose value is an odd number should have a left child leaf, otherwise it should have a right child leaf. For each internal node create a child leaf node at first, and then create a child internal node (if it exists). Assign values of the given sequence to tree nodes in order of their creation. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано четное число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать дерево из <i>N</i>&nbsp;вершин со следующей структурой: если вершина дерева является внутренней, то в случае, если она имеет нечетное значение, ее левая дочерняя вершина должна быть листом, а в случае, если она имеет четное значение, листом должна быть ее правая вершина. Для каждой внутренней вершины вначале создавать дочернюю вершину-лист, а затем дочернюю внутреннюю вершину (если данная вершина существует); каждой создаваемой вершине присваивать очередное значение из исходного набора. Вывести указатель на корень созданного дерева.

---

## Tree 30
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. Create a tree that satisfies the following conditions: the value of root node equals&nbsp;<i>N</i>; if the value of a node is an even number&nbsp;<i>K</i> then this node has only a left child whose value equals&nbsp;<i>K</i>/2; if the value of a node equals&nbsp;1 then this node is a leaf; if the value of a node is another odd number&nbsp;<i>K</i> then this node has a left child whose value equals&nbsp;<i>K</i>/2 and has a right child whose value equals&nbsp;<i>K</i>&nbsp;&#8722;&nbsp;<i>K</i>/2 (&#34;/&#34; denotes the operator of integer division). Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0). Создать дерево, корень которого имеет значение&nbsp;<i>N</i>, а вершины обладают следующими свойствами: вершина с четным значением&nbsp;<i>K</i> имеет левую дочернюю вершину со значением&nbsp;<i>K</i>/2 и не имеет правой дочерней вершины; вершина со значением 1 является листом; вершина с любым другим нечетным значением&nbsp;<i>K</i> имеет две дочерние вершины: левую со значением <i>K</i>/2 и правую со значением <i>K</i>&nbsp;&#8722;&nbsp;<i>K</i>/2 (символ &#171;/&#187; обозначает операцию деления нацело). Вывести указатель на корень созданного дерева.

---

## Tree 31
#### ![en](https://img.shields.io/badge/EN-blue) Two positive integers <i>L</i>, <i>N</i> (<i>N</i>&nbsp;&gt;&nbsp;<i>L</i>) and a sequence of <i>N</i>&nbsp;integers are given. Create a tree of depth&nbsp;<i>L</i>. Use elements of the given sequence as node values; add new nodes using the following algorithm: for each node of the level not greater than&nbsp;<i>L</i> create the node itself, then its left subtree of corresponding depth, and finally its right subtree. If less than&nbsp;<i>N</i> nodes are required to create an <i>L</i>-depth tree then do not use the rest of elements of the given sequence. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Даны положительные числа <i>L</i>, <i>N</i> (<i>N</i>&nbsp;&gt;&nbsp;<i>L</i>) и набор из <i>N</i>&nbsp;чисел. Создать дерево глубины&nbsp;<i>L</i>, содержащее вершины со значениями из исходного набора. Вершины добавлять к дереву в префиксном порядке, используя алгоритм, который для каждой вершины уровня, не превышающего&nbsp;<i>L</i>, вначале создает саму вершину с очередным значением из исходного набора, затем ее левое поддерево соответствующей глубины, а затем ее правое поддерево. Если для заполнения дерева глубины&nbsp;<i>L</i> требуется менее <i>N</i>&nbsp;вершин, то оставшиеся числа из исходного набора не использовать. Вывести указатель на корень созданного дерева.

---

## Tree 32
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Create <i>a balanced tree</i> with <i>N</i>&nbsp;nodes (that is, a binary tree which satisfies the following condition: for each tree node the amount of nodes of its left subtree differs at most on&nbsp;1 from the amount of nodes of its right subtree) and output the address of the tree root. Use elements of the given sequence as node values; create the tree by means of the following recursive algorithm: create a root node, then repeat the algorithm twice: for creating the left subtree with <i>N</i>/2&nbsp;nodes and for creating the right subtree with <i>N</i>&nbsp;&#8722;&nbsp;1&nbsp;&#8722;&nbsp;<i>N</i>/2 nodes (&#34;/&#34; denotes the operator of integer division).
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Создать <i>идеально сбалансированное дерево</i> из <i>N</i>&nbsp;вершин с заданными значениями (т.&nbsp;е. дерево, для каждой вершины которого количество вершин в его левом и правом поддереве отличается не более чем на&nbsp;1) и вывести указатель на его корень. Для создания дерева использовать рекурсивный алгоритм, который создает вершину дерева с очередным значением, после чего вызывается для создания левого поддерева с <i>N</i>/2&nbsp;вершинами и правого поддерева с <i>N</i>&nbsp;&#8722;&nbsp;1&nbsp;&#8722;&nbsp;<i>N</i>/2 вершинами (символ &#171;/&#187; обозначает операцию деления нацело).

---

## Tree 33
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. Create a balanced tree with <i>N</i>&nbsp;nodes and output the address of the tree root. The value of each node should be equal to its level (for example, the root value is&nbsp;0, the value of its children is&nbsp;1, and so&nbsp;on). Create the balanced tree by means of the recursive algorithm described in Tree32.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0). Создать идеально сбалансированное дерево из <i>N</i>&nbsp;вершин и вывести указатель на его корень. Значение каждой вершины положить равным уровню этой вершины (например, корень дерева должен иметь значение&nbsp;0, его дочерние вершины &#8212; значение&nbsp;1 и т.&nbsp;д.). При формировании дерева использовать алгоритм, описанный в задании Tree32.

---

## Tree 34
#### ![en](https://img.shields.io/badge/EN-blue) An address&nbsp;<i>P</i><sub>1</sub> of the root of a nonempty tree is given. Create a copy of the tree and output the address&nbsp;<i>P</i><sub>2</sub> of its root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Создать копию данного дерева и вывести указатель <i>P</i><sub>2</sub> на корень созданной копии.

---

## Tree 35
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Double the value of each tree node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Удвоить значение каждой вершины дерева.

---

## Tree 36
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Halve the value of each tree node whose initial value is an even number.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для каждой вершины дерева с четным значением уменьшить ее значение в два раза.

---

## Tree 37
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Add&nbsp;1 to the value of each tree leaf and subtract&nbsp;1 from the value of each internal node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Добавить&nbsp;1 к значению каждого листа дерева и вычесть&nbsp;1 из значения каждой внутренней вершины.

---

## Tree 38
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. For each tree node with two child swap values of its child nodes (that is, swap values of <i>Data</i> fields of child nodes).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для каждой вершины дерева, имеющей две дочерние вершины, поменять местами значения дочерних вершин (т.&nbsp;е. значения их полей Data).

---

## Tree 39
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Swap child nodes of each internal node in the tree (that is, swap values of its <i>Left</i> and <i>Right</i> field).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для всех внутренних вершин дерева поменять местами их левые и правые дочерние вершины (т.&nbsp;е. значения полей Left и Right).

---

## Tree 40
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Remove all nodes from the tree (except the root), release the memory allocated for removed nodes, and assign NULL to the <i>Left</i> and <i>Right</i> fields of the root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Удалить из дерева все вершины, кроме корня, и освободить память, которую занимали удаленные вершины (полям Left и Right корня следует присвоить значение NULL).

---

## Tree 41
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given, the tree contains at least two nodes. Remove all tree leaves and assign NULL to the <i>Left</i> and <i>Right</i> fields of their parents. Release the memory allocated for removed nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева, содержащего по крайней мере две вершины. Удалить каждую вершину дерева, являющуюся листом; при этом соответствующее поле родительской вершины (Left или Right) следует положить равным NULL. При удалении вершин освобождать память, которую они занимали.

---

## Tree 42
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Remove all nodes whose value is less than the root value, together with all their descendants. Release the memory allocated for removed nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Удалить вершины дерева, имеющие значения, меньшие значения корня, вместе со всеми их дочерними вершинами. При удалении вершин освобождать память, которую они занимали.

---

## Tree 43
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty tree is given. Apply the following action to each tree node that has two child nodes: if the node value is an even number then remove its right child, otherwise remove its left child. Use preorder tree walk; each node should be removed together with all its descendants. Release the memory allocated for removed nodes.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Для вершин дерева, имеющих две дочерние вершины, удалить одну из дочерних вершин: правую, если родительская вершина имеет четное значение, и левую в противном случае (вершины дерева перебирать в префиксном порядке, при удалении вершины удалять и всех ее потомков). Для удаленных вершин освобождать память, которую они занимали.

---

## Tree 44
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Add two child nodes to each tree leaf; the values of left and right child nodes should be equal to 10 and 11 respectively.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Ко всем вершинам дерева, которые являются листьями, добавить по две дочерние вершины-листа: левую со значением&nbsp;10 и правую со значением&nbsp;11.

---

## Tree 45
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Add one child node to each thee leaf; if the leaf value is an odd number then its child should be a left node, otherwise its child should be a right one. Value of created child node should be equal to value of its parent.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Ко всем вершинам дерева, которые являются листьями, добавить по одной дочерней вершине-листу; при этом к исходной вершине с нечетным значением добавляется левая дочерняя вершина, а к вершине с четным значением &#8212; правая. Значение каждой добавленной вершины положить равным значению ее родительской вершины.

---

## Tree 46
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. For each tree node with one child add another child node (a leaf). Value of created child node should be equal to doubled value of its parent.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Ко всем вершинам дерева, которые содержат ровно по одной дочерней вершине, добавить еще одну дочернюю вершину-лист. Значение каждой добавленной вершины положить равным удвоенному значению ее родительской вершины.

---

## Tree 47
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. Transform the given tree to a <i>perfect tree</i> by adding some new nodes (a perfect tree is a binary tree whose all leaves are at the same level). Do not change the initial depth of the tree; value of all new nodes should be equal to&nbsp;&#8722;1.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Не изменяя глубины&nbsp;<i>L</i> исходного дерева, дополнить его до <i>полного дерева</i>, т.&nbsp;е. дерева, у которого все внутренние вершины имеют по два непосредственных потомка, а все листья находятся на уровне&nbsp;<i>L</i>. Значения всех добавленных вершин положить равными&nbsp;&#8722;1.

---

## Tree 48
#### ![en](https://img.shields.io/badge/EN-blue) An address&nbsp;<i>P</i><sub>1</sub> of a tree node is given. Tree node is a record of TNode type containing the <i>Data</i> field (of integer type) and the <i>Left</i>, <i>Right</i>, and <i>Parent</i> fields (of PNode type). The <i>Left</i> and <i>Right</i> fields point to the left and right child nodes respectively, the <i>Parent</i> field points to the parent node (the <i>Parent</i> field of the root node equals NULL). Output pointers&nbsp;<i>P</i><sub><i>L</i></sub>, <i>P</i><sub><i>R</i></sub> to the left and right child of the given node, <i>P</i><sub>0</sub> to its parent, and&nbsp;<i>P</i><sub>2</sub> to its <i>sibling</i> (siblings are nodes that have the same parent). If some of required nodes are not exist then output NULL for each absent node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан адрес <i>P</i><sub>1</sub> вершины дерева &#8212; записи типа TNode, содержащей поля Data (целого типа), Left, Right и Parent (типа PNode &#8212; указателя на TNode). Поля Left и Right указывают на дочерние вершины, а поле Parent &#8212; на родительскую вершину данной вершины (если вершина является корнем дерева, то ее поле Parent равно NULL). Для данной вершины вывести указатели <i>P</i><sub><i>L</i></sub>, <i>P</i><sub><i>R</i></sub> и <i>P</i><sub>0</sub> на ее левую и правую дочерние вершины и родительскую вершину, а также указатель <i>P</i><sub>2</sub> на ее <i>сестру</i>, т.&nbsp;е. другую вершину дерева, имеющую в качестве родительской вершину с адресом <i>P</i><sub>0</sub>. Если некоторые из перечисленных вершин не существуют, то вывести для них значение NULL.

---

## Tree 49
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a tree is given. Tree nodes are represented by records of TNode type; they are linked by the <i>Left</i> and <i>Right</i> fields of TNode record. Using the <i>Parent</i> field of TNode record, transform the given tree into <i>a doubly linked tree</i> whose each node is connected not only with its child nodes (by the <i>Left</i> and <i>Right</i> fields) but also with its parent node (by the <i>Parent</i> field). The <i>Parent</i> field of the root node should be equal to NULL.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева, вершинами которого являются записи типа TNode, связанные между собой с помощью полей Left и Right. Используя поле Parent записи TNode, преобразовать исходное дерево в <i>дерево с обратной связью</i>, в котором каждая вершина связана не только со своими дочерними вершинами (полями Left и Right), но и с родительской вершиной (полем Parent). Поле Parent корня дерева положить равным NULL.

---

## Tree 50
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to some node of a doubly linked tree is given. Output the pointer&nbsp;<i>P</i><sub>2</sub> to the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на одну из вершин дерева с обратной связью. Вывести указатель <i>P</i><sub>2</sub> на корень исходного дерева.

---

## Tree 51
#### ![en](https://img.shields.io/badge/EN-blue) Pointers <i>P</i><sub>1</sub>, <i>P</i><sub>2</sub>, <i>P</i><sub>3</sub> to three nodes of a doubly linked tree are given. Output the level of each node (the level of the root equals&nbsp;0).
#### ![ru](https://img.shields.io/badge/RU-blue) Даны указатели <i>P</i><sub>1</sub>, <i>P</i><sub>2</sub>, <i>P</i><sub>3</sub> на три вершины дерева с обратной связью. Для каждой из данных вершин вывести ее уровень (корень дерева имеет уровень&nbsp;0).

---

## Tree 52
#### ![en](https://img.shields.io/badge/EN-blue) Pointers&nbsp;<i>P</i><sub>1</sub> and&nbsp;<i>P</i><sub>2</sub> to two different nodes of a doubly linked tree are given. Output <i>the degree of relationship</i> of the node&nbsp;<i>P</i><sub>1</sub> to the node&nbsp;<i>P</i><sub>2</sub>  (the degree of relationship equals &#8722;1 if the node&nbsp;<i>P</i><sub>2</sub> is not in the chain of ancestors of the node&nbsp;<i>P</i><sub>1</sub>; otherwise it equals <i>L</i><sub>1</sub>&nbsp;&#8722;&nbsp;<i>L</i><sub>2</sub>, where&nbsp;<i>L</i><sub>1</sub> and&nbsp;<i>L</i><sub>2</sub> are the levels of nodes&nbsp;<i>P</i><sub>1</sub> and&nbsp;<i>P</i><sub>2</sub> respectively).
#### ![ru](https://img.shields.io/badge/RU-blue) Даны указатели <i>P</i><sub>1</sub> и <i>P</i><sub>2</sub> на две различные вершины дерева с обратной связью. Вывести <i>степень родства</i> вершины <i>P</i><sub>1</sub> по отношению к вершине <i>P</i><sub>2</sub> (степень родства равна &#8722;1, если вершина <i>P</i><sub>2</sub> не находится в цепочке предков для вершины <i>P</i><sub>1</sub>; в противном случае степень родства равна <i>L</i><sub>1</sub>&nbsp;&#8722;&nbsp;<i>L</i><sub>2</sub>, где <i>L</i><sub>1</sub> и <i>L</i><sub>2</sub> &#8212; уровни вершин <i>P</i><sub>1</sub> и <i>P</i><sub>2</sub> соответственно).

---

## Tree 53
#### ![en](https://img.shields.io/badge/EN-blue) Pointers&nbsp;<i>P</i><sub>1</sub> and&nbsp;<i>P</i><sub>2</sub> to two different nodes of a doubly linked tree are given. Find the nearest mutual ancestor of the nodes&nbsp;<i>P</i><sub>1</sub> and&nbsp;<i>P</i><sub>2</sub> and output its pointer&nbsp;<i>P</i><sub>0</sub>.
#### ![ru](https://img.shields.io/badge/RU-blue) Даны указатели <i>P</i><sub>1</sub> и <i>P</i><sub>2</sub> на две различные вершины дерева с обратной связью. Вывести указатель <i>P</i><sub>0</sub> на вершину дерева, являющуюся ближайшим общим предком вершин <i>P</i><sub>1</sub> и <i>P</i><sub>2</sub>.

---

## Tree 54
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the node of a doubly linked tree is given. Create a copy of the given tree and output a pointer&nbsp;<i>P</i><sub>2</sub> to the root of the created tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на одну из вершин дерева с обратной связью. Создать копию данного дерева и вывести указатель <i>P</i><sub>2</sub> на корень созданной копии.

---

## Tree 55
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the non-root node of a doubly linked tree is given. If the node&nbsp;<i>P</i><sub>1</sub> has a sibling then remove the sibling together with all its descendants from the tree and release the memory allocated for removed nodes, if the node&nbsp;<i>P</i><sub>1</sub> has no sibling then create it and all its descendants as a copy of the subtree with the root&nbsp;<i>P</i><sub>1</sub>. Output the pointer&nbsp;<i>P</i><sub>0</sub> to the parent of&nbsp;<i>P</i><sub>1</sub>.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на вершину дерева с обратной связью, которая не является корнем. Если вершина <i>P</i><sub>1</sub> имеет сестру, то удалить эту сестру вместе со всеми ее потомками, освободив занимаемую ими память; если вершина <i>P</i><sub>1</sub> не имеет сестры, то создать сестру и всех ее потомков в виде копии поддерева с корнем <i>P</i><sub>1</sub>. Вывести указатель <i>P</i><sub>0</sub> на родительскую вершину вершины <i>P</i><sub>1</sub>.

---

## Tree 56
#### ![en](https://img.shields.io/badge/EN-blue) Two positive integers <i>L</i>, <i>N</i> (<i>N</i>&nbsp;&gt;&nbsp;<i>L</i>) and a sequence of <i>N</i>&nbsp;integers are given. Create a doubly linked tree of depth&nbsp;<i>L</i>. Use elements of the given sequence as node values; add new nodes using the following algorithm: for each node of the level not greater than&nbsp;<i>L</i> create the node itself, then its left subtree of corresponding depth, and finally its right subtree. If less than <i>N</i>&nbsp;nodes are required to create an <i>L</i>&nbsp;depth tree then do not use the rest of elements of the given sequence. Output the address of the tree root.
#### ![ru](https://img.shields.io/badge/RU-blue) Даны положительные числа <i>L</i>, <i>N</i> (<i>N</i>&nbsp;&gt;&nbsp;<i>L</i>) и набор из <i>N</i>&nbsp;чисел. Создать дерево глубины&nbsp;<i>L</i> с обратной связью, содержащее вершины со значениями из исходного набора. Вершины добавлять к дереву в префиксном порядке, используя алгоритм, который для каждой вершины уровня, не превышающего&nbsp;<i>L</i>, вначале создает саму вершину с очередным значением из исходного набора, затем ее левое поддерево соответствующей глубины, а затем ее правое поддерево. Если для заполнения дерева глубины&nbsp;<i>L</i> требуется менее <i>N</i>&nbsp;вершин, то оставшиеся числа из исходного набора не использовать. Вывести указатель на корень созданного дерева.

---

## Tree 57
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. It the tree is <i>a search tree</i>, that is, values of its nodes form a non-decreasing sequence in inorder tree walk, then output NULL; otherwise output the address of the first node (in inorder tree walk) that breaks the search-tree property.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Если данное дерево является <i>деревом поиска</i>, т.&nbsp;е. если при обходе его вершин в инфиксном порядке их значения образуют неубывающую последовательность, то вывести NULL; в противном случае вывести адрес первой вершины (в инфиксном порядке), нарушающей требуемую закономерность.

---

## Tree 58
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty tree is given. It the tree is <i>a non-recurrent search tree</i>, that is, values of its nodes form an increasing sequence in inorder tree walk, then output NULL; otherwise output the address of the first node (in inorder tree walk) that breaks the search-tree property.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Если данное дерево является <i>деревом поиска без повторяющихся элементов</i>, т.&nbsp;е. если при обходе его вершин в инфиксном порядке их значения образуют возрастающую последовательность, то вывести NULL; в противном случае вывести адрес первой вершины (в инфиксном порядке), нарушающей требуемую закономерность.

---

## Tree 59
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty non-recurrent search tree and an integer&nbsp;<i>K</i> are given. If the tree contains a node whose value equals&nbsp;<i>K</i> then output the address&nbsp;<i>P</i><sub>2</sub> of this node, otherwise output NULL. Also output the amount&nbsp;<i>N</i> of tree nodes that were checked during the search.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева поиска без повторяющихся элементов и число&nbsp;<i>K</i>. Определить, содержит ли исходное дерево вершину со значением&nbsp;<i>K</i>. Если содержит, то вывести указатель <i>P</i><sub>2</sub> на эту вершину, в противном случае вывести NULL. Вывести также количество&nbsp;<i>N</i> вершин, которые потребовалось проанализировать для выполнения задания.

---

## Tree 60
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty search tree and an integer&nbsp;<i>K</i> are given. Output the amount&nbsp;<i>C</i> of tree nodes whose value equals&nbsp;<i>K</i>. Also output the amount&nbsp;<i>N</i> of tree nodes that were checked during the search.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева поиска и число&nbsp;<i>K</i>. Вывести количество&nbsp;<i>C</i> вершин исходного дерева, имеющих значение&nbsp;<i>K</i>, а также количество&nbsp;<i>N</i> вершин, которые потребовалось проанализировать для выполнения задания.

---

## Tree 61
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a search tree and an integer&nbsp;<i>K</i> are given (if the tree is empty then <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Add a new node with the value&nbsp;<i>K</i> to the tree so that the tree still remains a search tree. Output the pointer <i>P</i><sub>2</sub> to the root of the resulting tree. Use the following recursive algorithm for a tree with the root&nbsp;<i>P</i>: if <i>P</i>&nbsp;=&nbsp;NULL then create a leaf with the value&nbsp;<i>K</i> and assign the address of the leaf to the pointer&nbsp;<i>P</i>; if the tree root exists then repeat the algorithm for the left subtree in case&nbsp;<i>K</i> is less than the root value or for the right subtree otherwise.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева поиска (если дерево является пустым, то <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Также дано число&nbsp;<i>K</i>. Добавить к исходному дереву поиска новую вершину со значением&nbsp;<i>K</i> таким образом, чтобы полученное дерево осталось деревом поиска, и вывести указатель <i>P</i><sub>2</sub> на корень полученного дерева. Использовать следующий рекурсивный алгоритм для дерева с корнем&nbsp;<i>P</i>: если <i>P</i>&nbsp;=&nbsp;NULL, то создать лист со значением&nbsp;<i>K</i> и присвоить указателю&nbsp;<i>P</i> адрес созданного листа; если корень&nbsp;<i>P</i> существует, то в случае, если его значение больше&nbsp;<i>K</i>, выполнить алгоритм для поля Left корня&nbsp;<i>P</i>, иначе выполнить алгоритм для его поля Right.

---

## Tree 62
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a non-recurrent search tree and an integer&nbsp;<i>K</i> are given (if the tree is empty then <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Add a new node with the value <i>K</i> to the tree so that the tree still remains a non-recurrent search tree. Do not change the given tree if it already contains a node with the value&nbsp;<i>K</i>. Output the pointer <i>P</i><sub>2</sub> to the root of the resulting tree. Use the following recursive algorithm for a tree with the root&nbsp;<i>P</i>: if <i>P</i>&nbsp;=&nbsp;NULL then create a leaf with the value&nbsp;<i>K</i> and assign the address of the leaf to the pointer&nbsp;<i>P</i>; if the tree root exists then repeat the algorithm for the left subtree in case&nbsp;<i>K</i> is less than the root value or for the right subtree in case&nbsp;<i>K</i> is greater than the root value.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень дерева поиска без повторяющихся элементов (если дерево является пустым, то <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Также дано число&nbsp;<i>K</i>. Добавить к исходному дереву поиска новую вершину со значением&nbsp;<i>K</i> таким образом, чтобы полученное дерево осталось деревом поиска без повторяющихся элементов, и вывести указатель&nbsp;<i>P</i><sub>2</sub> на корень полученного дерева. Если исходное дерево уже содержит вершину со значением&nbsp;<i>K</i>, то оставить дерево без изменений. Использовать следующий рекурсивный алгоритм для дерева с корнем&nbsp;<i>P</i>: если <i>P</i>&nbsp;=&nbsp;NULL, то создать лист со значением&nbsp;<i>K</i> и присвоить указателю&nbsp;<i>P</i> адрес созданного листа; если корень&nbsp;<i>P</i> существует, то в случае, если его значение больше&nbsp;<i>K</i>, выполнить алгоритм для поля Left корня&nbsp;<i>P</i>, а в случае, если его значение меньше&nbsp;<i>K</i>, выполнить алгоритм для его поля Right.

---

## Tree 63
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0), a sequence of <i>N</i>&nbsp;integers and a pointer <i>P</i><sub>1</sub> to the root of a search tree are given (if the tree is empty then <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Add <i>N</i>&nbsp;new nodes with values from the given sequence to the tree so that the tree still remains a search tree. Output the pointer&nbsp;<i>P</i><sub>2</sub> to the root of the resulting tree. Use the recursive algorithm described in Tree61 to add each new node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел, а также указатель <i>P</i><sub>1</sub> на корень дерева поиска (если дерево является пустым, то <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Добавить к исходному дереву поиска <i>N</i>&nbsp;новых вершин со значениями из исходного набора таким образом, чтобы полученное дерево осталось деревом поиска, и вывести указатель&nbsp;<i>P</i><sub>2</sub> на корень полученного дерева. Для добавления новых вершин использовать алгоритм, описанный в задании Tree61.

---

## Tree 64
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0), a sequence of <i>N</i>&nbsp;integers and a pointer <i>P</i><sub>1</sub> to the root of a non-recurrent search tree are given (if the tree is empty then <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Add <i>N</i>&nbsp;new nodes with values from the given sequence to the tree so that the tree still remains a non-recurrent search tree. Output the pointer&nbsp;<i>P</i><sub>2</sub> to the root of the resulting tree. Use the recursive algorithm described in Tree62 to add each new node.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел, а также указатель <i>P</i><sub>1</sub> на корень дерева поиска без повторяющихся элементов (если дерево является пустым, то <i>P</i><sub>1</sub>&nbsp;=&nbsp;NULL). Добавить к исходному дереву поиска новые вершины со значениями из исходного набора таким образом, чтобы полученное дерево осталось деревом поиска без повторяющихся элементов, и вывести указатель&nbsp;<i>P</i><sub>2</sub> на корень полученного дерева. Для добавления новых вершин использовать алгоритм, описанный в задании Tree62.

---

## Tree 65
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Sort the sequence by creating a search tree (use the recursive  algorithm described in Tree61 to add each new node). Output the pointer&nbsp;<i>P</i><sub>1</sub> to the root of the created tree. Also output elements of the sorted sequence using the inorder tree walk.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Отсортировать исходный набор чисел, создав для него дерево поиска (алгоритм добавления вершин к дереву поиска описан в задании Tree61). Вывести указатель&nbsp;<i>P</i><sub>1</sub> на корень полученного дерева, а также отсортированный набор чисел (для вывода набора чисел выполнить перебор вершин дерева в инфиксном порядке).

---

## Tree 66
#### ![en](https://img.shields.io/badge/EN-blue) An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a sequence of <i>N</i>&nbsp;integers are given. Sort all different elements of the sequence by creating a non-recurrent search tree (use the recursive algorithm described in Tree62 to add each new node). Output the pointer&nbsp;<i>P</i><sub>1</sub> to the root of the created tree. Also output elements of the sorted sequence using the inorder tree walk.
#### ![ru](https://img.shields.io/badge/RU-blue) Дано число&nbsp;<i>N</i> (&gt;&nbsp;0) и набор из <i>N</i>&nbsp;чисел. Получить отсортированный набор исходных чисел без повторений, создав для исходного набора дерево поиска без повторяющихся элементов (алгоритм добавления вершин к подобному дереву описан в задании Tree62). Вывести указатель <i>P</i><sub>1</sub> на корень полученного дерева, а также отсортированный набор чисел без повторений (для вывода набора чисел выполнить перебор вершин дерева в инфиксном порядке).

---

## Tree 67
#### ![en](https://img.shields.io/badge/EN-blue) Two pointers are given: <i>P</i><sub>1</sub> to the root of a nonempty search tree and&nbsp;<i>P</i><sub>2</sub> to one of its nodes with no more than one child. Remove the node&nbsp;<i>P</i><sub>2</sub> from the tree so that the tree still remains a search tree (if the node&nbsp;<i>P</i><sub>2</sub> has a child then link the child with the parent of the node&nbsp;<i>P</i><sub>2</sub>). If the resulting tree is not empty then output the pointer&nbsp;<i>P</i><sub>3</sub> to its root, otherwise output NULL.
#### ![ru](https://img.shields.io/badge/RU-blue) Даны два указателя: <i>P</i><sub>1</sub> на корень непустого дерева поиска и <i>P</i><sub>2</sub> на одну из вершин этого дерева, имеющих не более одной дочерней вершины. Удалить из исходного дерева вершину с адресом <i>P</i><sub>2</sub> так, чтобы полученное дерево осталось деревом поиска (если удаляемая вершина <i>P</i><sub>2</sub> имеет дочернюю вершину, то эту дочернюю вершину необходимо связать с родительской вершиной вершины <i>P</i><sub>2</sub>). Вывести указатель <i>P</i><sub>3</sub> на корень полученного дерева или NULL, если в результате удаления вершины <i>P</i><sub>2</sub> дерево стало пустым.

---

## Tree 68
#### ![en](https://img.shields.io/badge/EN-blue) Two pointers are given: <i>P</i><sub>1</sub> to the root of a nonempty search tree and&nbsp;<i>P</i><sub>2</sub> to one of its nodes with two children. Remove the node&nbsp;<i>P</i><sub>2</sub> from the tree so that the tree still remains a search tree. Use the following algorithm: find the node&nbsp;<i>P</i> with the maximal value in the left subtree of the node&nbsp;<i>P<sub>2</sub></i>, then assign its value to the node&nbsp;<i>P</i><sub>2</sub>, and finally remove the node&nbsp;<i>P</i> as in Tree67 (because the node&nbsp;<i>P</i> should have no more than one child).
#### ![ru](https://img.shields.io/badge/RU-blue) Даны два указателя: <i>P</i><sub>1</sub> на корень непустого дерева поиска и <i>P</i><sub>2</sub> на одну из вершин этого дерева, имеющих две дочерние вершины. Удалить из исходного дерева вершину <i>P</i><sub>2</sub> так, чтобы полученное дерево осталось деревом поиска. Удаление выполнять следующим образом: в левом поддереве вершины <i>P</i><sub>2</sub> найти вершину&nbsp;<i>P</i> с наибольшим значением, присвоить это наибольшее значение вершине <i>P</i><sub>2</sub>, после чего удалить вершину&nbsp;<i>P</i>, действуя, как в задании Tree67 (поскольку вершина&nbsp;<i>P</i> будет иметь не более одной дочерней вершины).

---

## Tree 69
#### ![en](https://img.shields.io/badge/EN-blue) Two pointers are given: <i>P</i><sub>1</sub> to the root of a nonempty search tree and&nbsp;<i>P</i><sub>2</sub> to one of its nodes with two children. Remove the node&nbsp;<i>P</i><sub>2</sub> from the tree so that the tree still remains a search tree. Use the following algorithm: find the node&nbsp;<i>P</i> with the minimal value in the right subtree of the node&nbsp;<i>P<sub>2</sub></i>, then assign its value to the node&nbsp;<i>P</i><sub>2</sub>, and finally remove the node&nbsp;<i>P</i> as in Tree67 (because the node&nbsp;<i>P</i> should have no more than one child).
#### ![ru](https://img.shields.io/badge/RU-blue) Даны два указателя: <i>P</i><sub>1</sub> на корень непустого дерева поиска и <i>P</i><sub>2</sub> на одну из вершин этого дерева, имеющих две дочерние вершины. Удалить из исходного дерева вершину <i>P</i><sub>2</sub> так, чтобы полученное дерево осталось деревом поиска. Удаление выполнять следующим образом: в правом поддереве вершины <i>P</i><sub>2</sub> найти вершину&nbsp;<i>P</i> с наименьшим значением, присвоить это наименьшее значение вершине <i>P</i><sub>2</sub>, после чего удалить вершину&nbsp;<i>P</i>, действуя, как в задании Tree67 (поскольку вершина&nbsp;<i>P</i> будет иметь не более одной дочерней вершины).

---

## Tree 70
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to a node of a doubly linked search tree is given. Remove the node&nbsp;<i>P</i><sub>1</sub> from the tree so that the tree still remains a doubly linked search tree. If the resulting tree is not empty then output the pointer&nbsp;<i>P</i><sub>2</sub> to its root, otherwise output NULL. If the node&nbsp;<i>P</i><sub>1</sub> has two children then use the algorithm described in Tree68 for its removing.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на одну из вершин непустого дерева поиска с обратной связью. Удалить из исходного дерева вершину <i>P</i><sub>1</sub> таким образом, чтобы полученное дерево осталось деревом поиска с обратной связью, и вывести указатель <i>P</i><sub>2</sub> на корень полученного дерева или NULL, если в результате удаления дерево стало пустым. Если вершина <i>P</i><sub>1</sub> имеет две дочерние вершины, то для ее удаления использовать алгоритм, описанный в задании Tree68.

---

## Tree 71
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to a node of a doubly linked search tree is given. Remove the node&nbsp;<i>P</i><sub>1</sub> from the tree so that the tree still remains a doubly linked search tree. If the resulting tree is not empty then output the pointer&nbsp;<i>P</i><sub>2</sub> to its root, otherwise output NULL. If the node&nbsp;<i>P</i><sub>1</sub> has two children then use the algorithm described in Tree69 for its removing.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на одну из вершин непустого дерева поиска с обратной связью. Удалить из исходного дерева вершину <i>P</i><sub>1</sub> таким образом, чтобы полученное дерево осталось деревом поиска с обратной связью, и вывести указатель <i>P</i><sub>2</sub> на корень полученного дерева или NULL, если в результате удаления дерево стало пустым. Если вершина <i>P</i><sub>1</sub> имеет две дочерние вершины, то для ее удаления использовать алгоритм, описанный в задании Tree69.

---

## Tree 72
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a nonempty tree is given. The tree representation is defined as follows (blank characters are not used): <table align="center"><tr><td align=left>&lt;tree&gt; <td align=center><tt>::=</tt>  <td align=left>&lt;empty string&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;node&gt;(&lt;left subtree&gt;,&lt;right subtree&gt;)<tr><td align=left>&lt;node&gt; <td align=center><tt>::=</tt>  <td align=left>&lt;digit&gt;</table>For example, &#34;4(2(,),6(,7(3(,),)))&#34;. Create a tree represented by the string&nbsp;<i>S</i> and output the pointer to its root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание непустого дерева в следующем формате: <table align="center"><tr><td align=left>&lt;дерево&gt;<td align=center><tt>::=</tt> <td align=left>&lt;пусто&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;вершина&gt;(&lt;левое поддерево&gt;,&lt;правое поддерево&gt;)<tr><td align=left>&lt;вершина&gt;<td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt;</table>Например, &#171;4(2(,),6(,7(3(,),)))&#187; (пробелы отсутствуют). Создать дерево по описанию, приведенному в&nbsp;<i>S</i>, и вывести указатель на его корень.

---

## Tree 73
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the string that describes the tree using the representation specified in Tree72.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести строку с описанием исходного дерева в формате, приведенном в задании Tree72.

---

## Tree 74
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a nonempty tree is given. The tree representation is defined as follows (blank characters are not used, the node representation depends on presence of subtrees of the node): <table align="center"><tr><td align=left>&lt;tree&gt; <td align=center><tt>::=</tt>  <td align=left>&lt;node&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;node&gt;(&lt;left subtree&gt;,&lt;right subtree&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;node&gt;(&lt;left subtree&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;node&gt;(,&lt;right subtree&gt;)<tr><td align=left>&lt;node&gt; <td align=center><tt>::=</tt>  <td align=left>&lt;digit&gt;</table>For example, &#34;4(2,6(,7(3)))&#34;. Create a tree represented by the string&nbsp;<i>S</i> and output the pointer to its root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание непустого дерева в следующем формате: <table align="center"><tr><td align=left>&lt;дерево&gt;<td align=center><tt>::=</tt> <td align=left>&lt;вершина&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;вершина&gt;(&lt;левое поддерево&gt;,&lt;правое поддерево&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;вершина&gt;(&lt;левое поддерево&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;вершина&gt;(,&lt;правое поддерево&gt;)<tr><td align=left>&lt;вершина&gt;<td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt;</table>Например, &#171;4(2,6(,7(3)))&#187; (пробелы отсутствуют, вид описания вершины зависит от того, имеет ли вершина непустое левое и/или правое поддерево). Создать дерево по описанию, приведенному в&nbsp;<i>S</i>, и вывести указатель на его корень.

---

## Tree 75
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty tree is given. Output the string that describes the tree using the representation specified in Tree74.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева. Вывести строку с описанием исходного дерева в формате, приведенном в задании Tree74.

---

## Tree 76
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a correct expression of integer type is given. The expression is defined as follows (blank characters are not used): <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>(&lt;expression&gt;&lt;operator&gt;&lt;expression&gt;)<tr><td align=left>&lt;operator&gt;   <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Create a tree that represents the given expression (<i>a parse tree</i>): each internal node corresponds to one of the arithmetic operators and equals &#8722;1 for addition, &#8722;2 for subtraction, and &#8722;3 for multiplication; a left subtree of a node-operator represents its left operand and a right subtree represents its right operand; leaf nodes represent digits. Output the pointer to the root of the created tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание числового выражения в следующем формате: <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>(&lt;выражение&gt;&lt;знак&gt;&lt;выражение&gt;)<tr><td align=left>&lt;знак&gt;      <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Пробелы в строке отсутствуют. Создать дерево, соответствующее исходному выражению (<i>дерево разбора выражения</i>): каждая внутренняя вершина дерева должна соответствовать одной из трех возможных арифметических операций и иметь значение &#8722;1 для операции сложения, &#8722;2 для операции вычитания и &#8722;3 для операции умножения; левое и правое дочерние поддеревья любой внутренней вершины-операции должны соответствовать выражениям слева и справа от знака операции; листьями полученного дерева должны быть выражения-цифры. Вывести указатель на корень созданного дерева.

---

## Tree 77
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a correct expression of integer type is given. The expression is defined as follows (<i>the parenthesis-free preorder format</i>): <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;operator&gt; &lt;expression&gt; &lt;expression&gt;<tr><td align=left>&lt;operator&gt;   <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Expressions are separated from each other and from the operators by one blank character. Create a parse tree for the given expression and output the pointer to its root. See the description of parse tree structure in Tree76; a left subtree of the node-operator corresponds to its first operand and a right subtree corresponds to its second operand.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание числового выражения в следующем формате (так называемый <i>префиксный бесскобочный формат</i> записи числового выражения): <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;знак&gt; &lt;выражение&gt; &lt;выражение&gt;<tr><td align=left>&lt;знак&gt;      <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Выражения отделяются друг от друга и от знаков операций ровно одним пробелом. Создать дерево разбора выражения и вывести указатель на его корень. Структура дерева разбора выражения описана в задании Tree76; для каждой вершины-операции ее левое поддерево должно соответствовать первому операнду данной операции, а правое поддерево &#8212; второму.

---

## Tree 78
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a correct expression of integer type is given. The expression is defined as follows (<i>the parenthesis-free postorder format</i>): <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;expression&gt; &lt;expression&gt; &lt;operator&gt;<tr><td align=left>&lt;operator&gt;   <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Expressions are separated from each other and from the operators by one blank character. Create a parse tree for the given expression and output the pointer to its root. See the description of parse tree structure in Tree76; a left subtree of the node-operator corresponds to its first operand and a right subtree corresponds to its second operand.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание числового выражения в следующем формате (так называемый <i>постфиксный бесскобочный формат</i> записи числового выражения): <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;выражение&gt; &lt;выражение&gt; &lt;знак&gt;<tr><td align=left>&lt;знак&gt;      <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table>Выражения отделяются друг от друга и от знаков операций ровно одним пробелом. Создать дерево разбора выражения и вывести указатель на его корень. Структура дерева разбора выражения описана в задании Tree76; для каждой вершины-операции ее левое поддерево должно соответствовать первому операнду данной операции, а правое поддерево &#8212; второму.

---

## Tree 79
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty parse tree is given (see the description of parse tree structure in Tree76). Output the value of expression that corresponds to the given tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения (структура дерева описана в задании Tree76). Вывести значение выражения, определяемого исходным деревом.

---

## Tree 80
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty parse tree is given (see the description of parse tree structure in Tree76). Output the string representation of expression that corresponds to the given tree. Use the expression format specified in the same task: <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>(&lt;expression&gt;&lt;operator&gt;&lt;expression&gt;)<tr><td align=left>&lt;operator&gt;   <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table></p>
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения (структура дерева описана в задании Tree76). Вывести строковое представление соответствующего выражения в формате, приведенном в том же задании: <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; |<tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>(&lt;выражение&gt;&lt;знак&gt;&lt;выражение&gt;)<tr><td align=left>&lt;знак&gt;      <td align=center><tt>::=</tt> <td align=left>+ | &#8722; | *</table></p>

---

## Tree 81
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty parse tree is given. Output the string representation of expression that corresponds to the given tree. Use the parenthesis-free preorder format (see Tree77).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения. Вывести строковое представление соответствующего выражения в префиксном бесскобочном формате (см. задание Tree77).

---

## Tree 82
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty parse tree is given. Output the string representation of expression that corresponds to the given tree. Use the parenthesis-free postorder format (see Tree78).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения. Вывести строковое представление соответствующего выражения в постфиксном бесскобочном формате (см. задание Tree78).

---

## Tree 83
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a correct expression of integer type is given. The expression is defined as follows (blank characters are not used, functions&nbsp;M and&nbsp;m return their maximal and minimal argument respectively): <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; | M(&lt;expression&gt;&nbsp;,&nbsp;&lt;expression&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>m(&lt;expression&gt;&nbsp;,&nbsp;&lt;expression&gt;)</table>Create a parse tree for the given expression: each internal node corresponds to one of two available functions and equals&nbsp;&#8722;1 for the function&nbsp;M and&nbsp;&#8722;2 for the function&nbsp;m; a left subtree of a node-function represents its first argument and a right subtree represents its second argument; leaf nodes represent digits. Output the pointer to the root of the created tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка <i>S</i>, содержащая описание числового выражения в следующем формате (функция&nbsp;M возвращает максимальное из двух выражений, а m &#8212; минимальное): <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; | M(&lt;выражение&gt;&nbsp;,&nbsp;&lt;выражение&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>m(&lt;выражение&gt;&nbsp;,&nbsp;&lt;выражение&gt;)</table>Пробелы в строке отсутствуют. Создать дерево разбора исходного выражения: каждая внутренняя вершина дерева должна соответствовать одной из двух возможных функций и иметь значение &#8722;1 для функции M и &#8722;2 для функции m; для каждой вершины-функции ее левое поддерево должно соответствовать первому аргументу функции, а правое поддерево &#8212; второму; листьями полученного дерева должны быть выражения-цифры. Вывести указатель на корень созданного дерева.

---

## Tree 84
#### ![en](https://img.shields.io/badge/EN-blue) A pointer <i>P</i><sub>1</sub> to the root of a nonempty parse tree is given (see the description of parse tree structure in Tree83). Output the value of expression that corresponds to the given tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения (структура дерева описана в задании Tree83). Вывести значение выражения, определяемого исходным деревом.

---

## Tree 85
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty parse tree is given (see the description of parse tree structure in Tree83). Output the string representation of expression that corresponds to the given tree. Use the expression format specified in the same task: <table align="center"><tr><td align=left>&lt;expression&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt; | M(&lt;expression&gt;&nbsp;,&nbsp;&lt;expression&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>m(&lt;expression&gt;&nbsp;,&nbsp;&lt;expression&gt;)</table></p>
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель <i>P</i><sub>1</sub> на корень непустого дерева разбора выражения (структура дерева описана в задании Tree83). Вывести строковое представление соответствующего выражения в формате, приведенном в том же задании: <table align="center"><tr><td align=left>&lt;выражение&gt; <td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt; | M(&lt;выражение&gt;&nbsp;,&nbsp;&lt;выражение&gt;) | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>m(&lt;выражение&gt;&nbsp;,&nbsp;&lt;выражение&gt;)</table></p>

---

## Tree 86
#### ![en](https://img.shields.io/badge/EN-blue) In a <i>general tree</i> a node may have more than two child nodes arranged in fixed order (from left to right). A general tree may be represented by linked records of the TNode type as follows: the <i>Left</i> field of any node points to its leftmost child whereas the <i>Right</i> field points to the nearest right sibling of this node. The tree root has no siblings, therefore its <i>Right</i> field always equals NULL. A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty binary tree is given. Create a general tree that corresponds to the given binary tree and output the pointer&nbsp;<i>P</i><sub>2</sub> to the root of the created general tree.
#### ![ru](https://img.shields.io/badge/RU-blue) <i>Дерево общего вида</i> (каждая вершина которого может иметь произвольное число дочерних вершин, расположенных в фиксированном порядке в направлении слева направо) реализуется с помощью набора связанных записей типа TNode следующим образом: для каждой внутренней вершины ее поле Left содержит указатель на ее первую (т.&nbsp;е. левую) дочернюю вершину, а поле Right&nbsp;&#8212; указатель на ее правую <i>сестру</i>, т.&nbsp;е. вершину, имеющую в дереве общего вида того же родителя. Поле Right корня дерева общего вида всегда равно NULL, так как корень сестер не имеет. Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого бинарного дерева. Создать дерево общего вида, соответствующее исходному бинарному дереву, и вывести указатель&nbsp;<i>P</i><sub>2</sub> на его корень.

---

## Tree 87
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Each node has no more than two child nodes. Create a binary tree corresponding to the given general tree and output the pointer&nbsp;<i>P</i><sub>2</sub> to the root of the created binary tree. First child of any node of general tree should be the left child of the correspondent node of binary tree.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Любая вершина исходного дерева имеет не более двух дочерних вершин. Создать бинарное дерево, соответствующее исходному дереву общего вида, и вывести указатель&nbsp;<i>P</i><sub>2</sub> на его корень. Считать, что первая дочерняя вершина любой вершины дерева общего вида соответствует левой дочерней вершине в бинарном дереве.

---

## Tree 88
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output <i>the depth</i> of the tree (that is, the maximal level of tree nodes). All siblings are assumed to be on the same level; the level of the root equals&nbsp;0.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести <i>глубину</i> данного дерева, т.&nbsp;е. значение его максимального уровня, считая, что все вершины-сестры находятся на одном уровне, а корень дерева расположен на уровне&nbsp;0.

---

## Tree 89
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. For each tree level (including the zero one) output the amount of its nodes. The tree depth is assumed to be not greater than&nbsp;10.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Для каждого из уровней данного дерева, начиная с уровня&nbsp;0, вывести количество вершин, находящихся на этом уровне. Считать, что глубина дерева общего вида не превосходит&nbsp;10.

---

## Tree 90
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. For each tree level (including the zero one) output the sum of values of its nodes. The tree depth is assumed to be not greater than&nbsp;10.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Для каждого из уровней данного дерева, начиная с уровня&nbsp;0, вывести сумму значений вершин, находящихся на этом уровне. Считать, что глубина дерева общего вида не превосходит&nbsp;10.

---

## Tree 91
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree and an integer&nbsp;<i>L</i> (&#8805;&nbsp;0) are given. Output the values of nodes of the level&nbsp;<i>L</i> and their amount&nbsp;<i>N</i> (nodes must be ordered from left to right). If nodes of the level&nbsp;<i>L</i> are absent then output&nbsp;0.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Также дано неотрицательное число&nbsp;<i>L</i>. Перебирая дочерние вершины в заданном порядке (т.&nbsp;е. слева направо), вывести значения всех вершин уровня&nbsp;<i>L</i> и их количество&nbsp;<i>N</i>. Если дерево не содержит вершин уровня&nbsp;<i>L</i>, то вывести&nbsp;0.

---

## Tree 92
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output the values of all nodes using <i>inorder tree walk</i> as follows: output the first (i.&nbsp;e., the most left) subtree using inorder tree walk, then output the root value, then output all other subtrees (from left to right) using inorder tree walk too.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести значения всех вершин дерева в <i>инфиксном порядке</i>: вначале выводится содержимое первого (левого) поддерева в инфиксном порядке, затем выводится значение корня, а затем &#8212; содержимое остальных поддеревьев в инфиксном порядке (поддеревья перебираются слева направо).

---

## Tree 93
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output the values of all nodes using <i>postorder tree walk</i> as follows: output all subtrees (from left to right) using postorder tree walk, then output the root value.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести значения всех вершин дерева в <i>постфиксном порядке</i>: вначале выводится содержимое каждого поддерева в постфиксном порядке (поддеревья перебираются слева направо), а затем&nbsp;&#8212; значение корня.

---

## Tree 94
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree and an integer&nbsp;<i>N</i> (&#8805;&nbsp;0) are given. Output the amount of nodes that have exactly <i>N</i>&nbsp;child nodes (this amount may be equal to&nbsp;0).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Также дано неотрицательное число&nbsp;<i>N</i>. Вывести количество вершин исходного дерева, имеющих ровно <i>N</i>&nbsp;дочерних вершин. Если требуемые вершины отсутствуют, то вывести&nbsp;0.

---

## Tree 95
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output the pointer&nbsp;<i>P</i><sub>2</sub> to the first node that has the maximal amount of child nodes. Use inorder tree walk (see Tree92).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести указатель&nbsp;<i>P</i><sub>2</sub> на первую вершину дерева, имеющую наибольшее количество дочерних вершин. Вершины перебирать в инфиксном порядке (см. задание Tree92).

---

## Tree 96
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output the pointer&nbsp;<i>P</i><sub>2</sub> to the last node that has the maximal sum of values of child nodes. Use postorder tree walk (see Tree93).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести указатель&nbsp;<i>P</i><sub>2</sub> на последнюю вершину дерева с наибольшей суммой значений дочерних вершин. Вершины перебирать в постфиксном порядке (см. задание Tree93).

---

## Tree 97
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. In each set of siblings find the maximal value (that is, the maximal <i>Data</i> field) and assign this value to each node of the set.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. В каждом наборе вершин-сестер заменить все значения вершин (т.&nbsp;е. значения полей Data) на максимальное из их исходных значений.

---

## Tree 98
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. In each set of siblings inverse the order of the node values (that is, swap the <i>Data</i> fields of the first and the last sibling, then swap the <i>Data</i> fields of the second and the last but one sibling, and so&nbsp;on).
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. В каждом наборе вершин-сестер изменить порядок следования их значений на противоположный, т.&nbsp;е. поменять местами значения поля Data первой (левой) и последней (правой) сестры, второй и предпоследней сестры, и т.&nbsp;д.

---

## Tree 99
#### ![en](https://img.shields.io/badge/EN-blue) A string&nbsp;<i>S</i> that represents a nonempty general tree is given. The tree representation is defined as follows (blank characters are not used, siblings are ordered from left to right): <table align="center"><tr><td align=left>&lt;tree&gt; <td align=center><tt>::=</tt> <td align=left>&lt;node&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;node&gt;(&lt;list of subtrees&gt;)<tr><td align=left>&lt;list of subtrees&gt; <td align=center><tt>::=</tt> <td align=left>&lt;tree&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;tree&gt;,&lt;list of subtrees&gt;<tr><td align=left>&lt;node&gt; <td align=center><tt>::=</tt> <td align=left>&lt;digit&gt;</table>For example, &#34;3(2,7(6,4,5),8(4(2,3),5(1)))&#34;. Create a tree represented by the string&nbsp;<i>S</i> and output the pointer to its root.
#### ![ru](https://img.shields.io/badge/RU-blue) Дана строка&nbsp;<i>S</i>, содержащая описание непустого дерева общего вида в следующем формате: <table align="center"><tr><td align=left>&lt;дерево&gt;<td align=center><tt>::=</tt> <td align=left>&lt;вершина&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;вершина&gt;(&lt;список поддеревьев&gt;)<tr><td align=left>&lt;список поддеревьев&gt;<td align=center><tt>::=</tt> <td align=left>&lt;дерево&gt; | <tr><td align=left>&nbsp;<td align=center>&nbsp;<td align=left>&lt;дерево&gt;,&lt;список поддеревьев&gt;<tr><td align=left>&lt;вершина&gt;<td align=center><tt>::=</tt> <td align=left>&lt;цифра&gt;</table>Например, &#171;3(2,7(6,4,5),8(4(2,3),5(1)))&#187; (пробелы отсутствуют, вершины-сестры перечисляются в порядке слева направо). Создать дерево общего вида по описанию, приведенному в&nbsp;<i>S</i>, и вывести указатель на его корень.

---

## Tree 100
#### ![en](https://img.shields.io/badge/EN-blue) A pointer&nbsp;<i>P</i><sub>1</sub> to the root of a nonempty general tree is given. Output the string that describes the tree using the representation specified in Tree99.
#### ![ru](https://img.shields.io/badge/RU-blue) Дан указатель&nbsp;<i>P</i><sub>1</sub> на корень непустого дерева общего вида. Вывести строку с описанием исходного дерева в формате, приведенном в задании Tree99.

---

