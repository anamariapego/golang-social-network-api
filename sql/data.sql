-- DADOS DA TABELA DE USUÁRIOS

-- password_user: 12345
INSERT INTO users (name_user, nick, email, password_user)
VALUES
    ('Alice Silva', 'alice_s', 'alice.silva@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI'),
    ('Bruno Costa', 'bruno_c', 'bruno.costa@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI'),
    ('Carla Mendes', 'carla_m', 'carla.mendes@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI'),
    ('Diego Lima', 'diego_l', 'diego.lima@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI'),
    ('Tiago Torres', 'tiago_t', 'tiago.torres@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI'),
    ('Fernando Almeida', 'fernando_a', 'fernando.almeida@gmail.com', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzA2NzYyOTIsInVzZXJJZCI6MTB9.hv7flZbxTvXkTWfn1tPf8Nak8nnRkqfyf4dK_jApMHI');


-- DADOS DA TABELA DE SEGUIDORES

INSERT INTO followers (user_id, follower_id)
VALUES
    (1, 2), -- Alice segue Bruno
    (1, 3), -- Alice segue Carla
    (1, 3), -- Alice segue Fernando
    (2, 1), -- Bruno segue Alice
    (2, 4), -- Bruno segue Diego
    (3, 1), -- Carla segue Alice
    (3, 5), -- Carla segue Tiago
    (4, 6), -- Diego segue Fernando
    (4, 2), -- Diego segue Bruno
    (5, 3), -- Tiago segue Carla
    (6, 1), -- Fernando segue Alice
    (6, 5); -- Fernando segue Tiago


-- DADOS DA TABELA DE PUBLICAÇÕES

INSERT INTO publications (title, text, author_id)
VALUES
    ('Primeira publicação de Alice', 'Hoje estou começando minha jornada no mundo das publicações. Estou muito empolgada para compartilhar minhas ideias e experiências com todos vocês. Sempre gostei de escrever, e este é o momento perfeito para começar a compartilhar minhas histórias. Nos próximos posts, espero poder falar sobre viagens, aprendizado pessoal e como encontrei equilíbrio entre trabalho e vida pessoal. Estou muito animada para conectar com todos vocês!', 1),
    ('Publicação de Bruno', 'A tecnologia tem evoluído de forma impressionante, e com ela, surgem novas oportunidades e desafios. Em minha última experiência, participei de um projeto que envolvia inteligência artificial e automação, e fiquei impressionado com o impacto que essas tecnologias podem ter em diferentes setores da economia. Acredito que o futuro da tecnologia é brilhante, mas é preciso estarmos preparados para lidar com as mudanças que ela traz. Vamos discutir mais sobre as inovações tecnológicas e como elas podem transformar o nosso dia a dia.', 2),
    ('Pensamentos de Carla', 'Sempre fui apaixonada por viagens. Cada lugar que visito, cada cultura que conheço, me ensina algo novo sobre o mundo e sobre mim mesma. Uma das minhas viagens mais recentes foi para o Japão, um país que admiro pela sua rica história e pela fusão de tradições antigas com a modernidade. Durante a viagem, percebi como a convivência com culturas diferentes pode expandir nossa visão de mundo e nos tornar mais empáticos. Em breve, vou compartilhar mais detalhes sobre essa experiência, incluindo dicas de viagem e lugares imperdíveis.', 3),
    ('Diego e suas aventuras', 'Viajar é uma das minhas grandes paixões, e no último mês, tive a oportunidade de explorar algumas das regiões mais remotas do Brasil. Desde as cachoeiras de Goiás até as praias desertas do Maranhão, cada destino foi uma nova descoberta. Além das paisagens deslumbrantes, encontrei pessoas incríveis que me mostraram o verdadeiro significado de hospitalidade. Para quem gosta de aventura e natureza, o Brasil tem muito a oferecer. Em breve, vou escrever sobre os lugares mais incríveis que visitei e como se planejar para uma viagem como essa.', 4),
    ('Reflexões de Tiago', 'Nos últimos tempos, tenho refletido muito sobre como o estilo de vida que escolhemos pode impactar nossa saúde e bem-estar. Acredito que a chave para uma vida saudável está no equilíbrio: entre o trabalho e o lazer, entre a alimentação saudável e o prazer de saborear algo indulgente de vez em quando. A prática de exercícios físicos também tem um papel fundamental, não só para manter o corpo em forma, mas para aliviar o estresse do dia a dia. Nos próximos posts, quero compartilhar algumas dicas sobre como incorporar hábitos saudáveis no nosso cotidiano e como fazer essas mudanças de forma gradual, sem pressa.', 5),
    ('Fernando sobre o mercado de trabalho', 'O mercado de trabalho tem passado por transformações significativas nos últimos anos, impulsionadas pela tecnologia e pelas mudanças nas necessidades do consumidor. O trabalho remoto, por exemplo, se tornou uma realidade para muitas empresas, mas com isso surgem novas demandas em termos de produtividade, equilíbrio entre vida profissional e pessoal e gestão de equipes à distância. Acredito que o futuro do trabalho está cada vez mais relacionado à flexibilidade, mas também à capacidade de adaptação. Quero compartilhar algumas das tendências mais importantes do mercado e como podemos nos preparar para o futuro do trabalho.', 6);
    