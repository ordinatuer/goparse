# goparse
<h3>UNLOGGED</h3>
<code>ALTER TABLE "table" SET UNLOGGED</code><br>
2m18.193502895s против 4m22.809218056s без этого<br>
Естественно, есть нюансы<br>
<code>ALTER TABLE "table" SET LOGGED</code><br><hr>
<h3>Горутины</h3>
Параллельная работа со всеми файлами одновременно<br>
2m20.511416205s - с логированием и индексом
2m9.654263585s - без логирования
