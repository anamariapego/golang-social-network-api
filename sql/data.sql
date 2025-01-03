-- DADOS DA TABELA DE USUÁRIOS

-- password_user: 1234
INSERT INTO users (name_user, nick, email, password_user)
VALUES
    ('Alice Silva', 'alice_s', 'alice.silva@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.'),
    ('Bruno Costa', 'bruno_c', 'bruno.costa@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.'),
    ('Carla Mendes', 'carla_m', 'carla.mendes@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.'),
    ('Diego Lima', 'diego_l', 'diego.lima@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.'),
    ('Tiago Torres', 'tiago_t', 'tiago.torres@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.'),
    ('Fernando Almeida', 'fernando_a', 'fernando.almeida@gmail.com', '$2a$10$I7x/LlBwwNkOjYvdNuVKjeOoNjdGypF/RLQbfBNrrWuPbfutY7TX.');


-- DADOS DA TABELA DE SEGUIDORES

INSERT INTO followers (user_id, follower_id)
VALUES
    (1, 2), -- Alice segue Bruno
    (1, 3), -- Alice segue Carla
    (1, 6), -- Alice segue Fernando
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
    ('Primeira publicação de Alice', 'Hoje estou começando minha jornada no mundo das publicações. Estou muito empolgada para compartilhar minhas ideias e experiências com todos vocês. Sempre gostei de escrever, e este é o momento perfeito para começar a compartilhar minhas histórias. Espero poder falar sobre viagens, aprendizado pessoal e como encontrei equilíbrio entre trabalho e vida pessoal. Estou animada para conectar com todos!', 1),
    ('Publicação de Bruno', 'A tecnologia tem evoluído de forma impressionante, e com ela, surgem novas oportunidades e desafios. Participei de um projeto envolvendo inteligência artificial e automação. Fiquei impressionado com o impacto dessas tecnologias em diversos setores. Acredito que o futuro da tecnologia é brilhante, mas é preciso estarmos preparados para lidar com as mudanças. Vamos discutir mais sobre inovações tecnológicas.', 2),
    ('Pensamentos de Carla', 'Sempre fui apaixonada por viagens. Cada lugar que visito me ensina algo novo sobre o mundo e sobre mim mesma. Recentemente viajei para o Japão, um país que admiro pela sua rica história. Durante a viagem, percebi como a convivência com culturas diferentes pode expandir nossa visão de mundo. Em breve, vou compartilhar mais detalhes dessa experiência, incluindo dicas de viagem e lugares imperdíveis.', 3),
    ('Diego e suas aventuras', 'Viajar é uma das minhas grandes paixões. No último mês, explorei algumas das regiões mais remotas do Brasil, desde as cachoeiras de Goiás até as praias desertas do Maranhão. Cada destino foi uma nova descoberta, com paisagens deslumbrantes e pessoas incríveis que me mostraram o verdadeiro significado de hospitalidade. O Brasil tem muito a oferecer para quem gosta de aventura e natureza.', 4),
    ('Reflexões de Tiago', 'Nos últimos tempos, tenho refletido muito sobre o estilo de vida que escolhemos e como ele impacta nossa saúde e bem-estar. Acredito que a chave para uma vida saudável está no equilíbrio, entre o trabalho e o lazer, e na prática de exercícios físicos. Quero compartilhar algumas dicas sobre como incorporar hábitos saudáveis no nosso cotidiano e como fazer essas mudanças de forma gradual, sem pressa.', 5),
    ('Fernando sobre o mercado de trabalho', 'O mercado de trabalho tem passado por transformações nos últimos anos, impulsionadas pela tecnologia e pelas mudanças nas necessidades do consumidor. O trabalho remoto se tornou uma realidade, mas com isso surgem novas demandas em produtividade, equilíbrio entre vida profissional e pessoal. Acredito que o futuro do trabalho está cada vez mais relacionado à flexibilidade e à capacidade de adaptação.', 6);