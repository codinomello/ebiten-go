package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game{} -> representa o estado do jogo
type Game struct {
	leader *Leader
	florim []*Florim
}

// Sprite{} -> representa uma imagem e sua localização no jogo
type Sprite struct {
	x, y    float64
	image   *ebiten.Image
	display bool
}

// Leader{} -> representa o jogador principal
type Leader struct {
	sprite *Sprite
	name   string
}

// Florim{} -> representa uma criatura que segue o líder
type Florim struct {
	sprite *Sprite
}

// NewGame() -> cria um novo estado do jogo
func NewGame() *Game {
	florim := []*Florim{}
	for i := 0; i < 5; i++ {
		florim = append(florim, &Florim{sprite: &Sprite{x: 100 + float64(i*30), y: 100}})
	}

	// Definir um nome fixo para o líder
	leaderName := "Leader1"

	return &Game{
		leader: &Leader{
			sprite: &Sprite{
				x: 100.0,
				y: 100.0,
			},
			name: leaderName, // agora usando o nome fixo
		},
		florim: florim,
	}
}

// Update() -> atualiza o estado do jogo
func (g *Game) Update() error {
	// Movimenta o líder com as teclas de setas
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.leader.sprite.y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.leader.sprite.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.leader.sprite.y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.leader.sprite.x += 2
	}

	// Atualiza os florins para seguirem o líder
	for i, florim := range g.florim {
		targetX, targetY := g.leader.sprite.x, g.leader.sprite.y
		if i > 0 {
			// Os florins seguem o florim anterior
			targetX, targetY = g.florim[i-1].sprite.x, g.florim[i-1].sprite.y
		}

		// Calcula a distância e move gradualmente
		dx, dy := targetX-florim.sprite.x, targetY-florim.sprite.y
		distance := math.Hypot(dx, dy)
		if distance > 20 {
			florim.sprite.x += dx / distance * 2
			florim.sprite.y += dy / distance * 2
		}
	}

	return nil
}

// Draw() -> desenha o estado do jogo
func (g *Game) Draw(screen *ebiten.Image) {
	// Cor branca
	white := color.RGBA{255, 255, 255, 255}

	// Desenha o líder como um retângulo branco
	vector.DrawFilledRect(screen, float32(g.leader.sprite.x-10), float32(g.leader.sprite.y-10), 20, 20, white, false)

	// Desenha os Florins como retângulos menores
	for _, florim := range g.florim {
		vector.DrawFilledRect(screen, float32(florim.sprite.x-7.5), float32(florim.sprite.y-7.5), 15, 15, white, false)
	}
}

// Layout define o tamanho da tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

// Função principal
func main() {
	// Cria o estado do jogo
	game := NewGame()

	// Configura o título da janela e o tamanho
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Florins - Liderando o Caminho!")

	// Executa o jogo
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
